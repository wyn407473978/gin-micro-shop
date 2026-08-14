package discovery

import (
	"context"
	"encoding/json"
	"errors"
	registry2 "gin-micro-shop/pkg/etcd/registry"
	"log"
	"strings"
	"sync"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Discovery struct {
	client     *clientv3.Client
	serviceMap map[string][]*registry2.ServiceInstance
	lock       sync.RWMutex
}

func NewDiscovery(client *clientv3.Client) *Discovery {
	return &Discovery{
		client:     client,
		serviceMap: make(map[string][]*registry2.ServiceInstance),
	}
}

func (d *Discovery) WatchService(serviceName string) error {
	prefix := "/services/" + serviceName + "/"

	// 初始获取所有实例
	resp, err := d.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	d.updateServiceInstances(serviceName, resp.Kvs)

	// 监听变化
	go d.watch(prefix, serviceName)
	return nil
}

func (d *Discovery) watch(prefix, serviceName string) {
	watchChan := d.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for resp := range watchChan {
		for _, ev := range resp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				instance := parseInstance(ev.Kv)
				d.addInstance(serviceName, instance)
			case mvccpb.DELETE:
				instanceID := extractInstanceID(string(ev.Kv.Key))
				d.removeInstance(serviceName, instanceID)
			}
		}
	}
}

func (d *Discovery) updateServiceInstances(serviceName string, kvs []*mvccpb.KeyValue) {
	instances := make([]*registry2.ServiceInstance, 0, len(kvs))
	for _, kv := range kvs {
		instance := parseInstance(kv)
		instances = append(instances, instance)
	}

	d.lock.Lock()
	d.serviceMap[serviceName] = instances
	d.lock.Unlock()
}

func (d *Discovery) addInstance(serviceName string, instance *registry2.ServiceInstance) {
	d.lock.Lock()
	defer d.lock.Unlock()
	log.Printf("service instance added: service=%s, id=%s", serviceName, instance.ID)
	instances := d.serviceMap[serviceName]
	for _, ins := range instances {
		if ins.ID == instance.ID {
			return // 已存在
		}
	}
	d.serviceMap[serviceName] = append(instances, instance)
}

func (d *Discovery) removeInstance(serviceName, instanceID string) {
	d.lock.Lock()
	defer d.lock.Unlock()
	log.Printf("service instance removed: service=%s, id=%s", serviceName, instanceID)
	instances := d.serviceMap[serviceName]
	for i, ins := range instances {
		if ins.ID == instanceID {
			d.serviceMap[serviceName] = append(instances[:i], instances[i+1:]...)
			return
		}
	}
}

func (d *Discovery) GetInstances(serviceName string) ([]*registry2.ServiceInstance, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	instances, ok := d.serviceMap[serviceName]
	if !ok {
		return nil, errors.New("service not found")
	}
	return instances, nil
}

func parseInstance(kv *mvccpb.KeyValue) *registry2.ServiceInstance {
	instance := &registry2.ServiceInstance{}
	json.Unmarshal(kv.Value, instance)
	return instance
}

func extractInstanceID(key string) string {
	// 从 /services/user-service/192.168.1.10:8080 中提取 192.168.1.10:8080
	return key[strings.LastIndex(key, "/")+1:]
}
