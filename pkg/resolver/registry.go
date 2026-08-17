package resolver

import (
	"context"
	"encoding/json"
	"log"
	_ "time"

	"go.etcd.io/etcd/client/v3"
)

type ServiceInstance struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Address  string            `json:"address"`
	Port     int               `json:"port"`
	Metadata map[string]string `json:"metadata"`
}

type EtcdRegistry struct {
	client   *clientv3.Client
	leaseID  clientv3.LeaseID
	key      string
	ttl      int64
	stopChan chan struct{}
}

func NewEtcdRegistry(client *clientv3.Client, ttl int64) *EtcdRegistry {
	return &EtcdRegistry{
		client:   client,
		ttl:      ttl,
		stopChan: make(chan struct{}),
	}
}

func (r *EtcdRegistry) Register(instance *ServiceInstance) error {
	// 1. 创建租约
	leaseResp, err := r.client.Grant(context.Background(), r.ttl)
	if err != nil {
		return err
	}
	r.leaseID = leaseResp.ID

	// 2. 构造键值
	r.key = "/services/" + instance.Name + "/" + instance.ID
	value, err := json.Marshal(instance)
	if err != nil {
		return err
	}

	// 3. 注册服务
	_, err = r.client.Put(context.Background(), r.key, string(value), clientv3.WithLease(r.leaseID))
	if err != nil {
		return err
	}

	// 4. 保持心跳
	go r.keepAlive()

	return nil
}

func (r *EtcdRegistry) keepAlive() {
	keepAliveChan, err := r.client.KeepAlive(context.Background(), r.leaseID)
	if err != nil {
		log.Printf("keep alive failed: %v", err)
		return
	}

	for {
		select {
		case <-r.stopChan:
			return
		case _, ok := <-keepAliveChan:
			if !ok {
				log.Println("keep alive channel closed")
				go r.keepAlive() // 重新建立keepalive
				return
			}
		}
	}
}

func (r *EtcdRegistry) Deregister() error {
	close(r.stopChan)
	_, err := r.client.Revoke(context.Background(), r.leaseID)
	if err != nil {
		return err
	}
	_, err = r.client.Delete(context.Background(), r.key)
	return err
}
