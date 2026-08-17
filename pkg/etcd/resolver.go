package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
)

const Scheme = "etcd"

type EtcdResolverBuilder struct {
	client *clientv3.Client
	prefix string
}

func NewEtcdResolverBuilder(client *clientv3.Client, prefix string) *EtcdResolverBuilder {

	return &EtcdResolverBuilder{
		client: client,
		prefix: prefix,
	}
}

func (b *EtcdResolverBuilder) Scheme() string {
	return Scheme
}

func (b *EtcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	//如果我们 grpc.NewClient("etcd://user")  这里的到的serviceName 就是user
	serviceName := target.Endpoint()

	ctx, cancel := context.WithCancel(context.Background())
	r := &EtcdResolver{
		client:      b.client,
		cc:          cc,
		serviceName: serviceName,
		prefix:      b.prefix,
		ctx:         ctx,
		cancel:      cancel,
	}

	// 第一次立即查询
	if err := r.resolve(); err != nil {
		cancel()
		return nil, err
	}

	// 后台监听 etcd 服务变化
	go r.watch()

	return r, nil
}

type EtcdResolver struct {
	client *clientv3.Client

	// gRPC 给我们的
	cc resolver.ClientConn

	serviceName string
	prefix      string

	ctx    context.Context
	cancel context.CancelFunc
}

// 获取制定前缀的相关服务,更新到state中去
func (r *EtcdResolver) resolve() error {

	prefix := fmt.Sprintf(
		"%s/%s/",
		r.prefix,
		r.serviceName,
	)

	resp, err := r.client.Get(
		r.ctx,
		prefix,
		clientv3.WithPrefix(),
	)
	if err != nil {
		return err
	}

	addresses := make(
		[]resolver.Address,
		0,
		len(resp.Kvs),
	)

	for _, kv := range resp.Kvs {

		instance := &ServiceInstance{}
		err := json.Unmarshal(kv.Value, instance)
		if err != nil {
			return err
		}

		addresses = append(
			addresses,
			resolver.Address{
				Addr: net.JoinHostPort(instance.Address, fmt.Sprint(instance.Port)),
			},
		)
	}

	return r.cc.UpdateState(
		resolver.State{
			Addresses: addresses,
		},
	)
}

// 监控制定服务的列表变化 如果有变化就调用resolve 接口重新获取完整服务列表
func (r *EtcdResolver) watch() {

	prefix := fmt.Sprintf(
		"%s/%s/",
		r.prefix,
		r.serviceName,
	)

	watchChan := r.client.Watch(
		r.ctx,
		prefix,
		clientv3.WithPrefix(),
	)

	for watchResp := range watchChan {

		if err := watchResp.Err(); err != nil {
			r.cc.ReportError(err)
			continue
		}

		if len(watchResp.Events) == 0 {
			continue
		}

		// 发现服务发生变化
		// 重新获取完整服务列表
		log.Fatalf("service changed")
		if err := r.resolve(); err != nil {
			r.cc.ReportError(err)
		}
	}
}

func (r *EtcdResolver) ResolveNow(
	resolver.ResolveNowOptions,
) {
	// 我们已经通过 etcd Watch 主动接收变化，
	// 所以这里暂时不用处理。
}

func (r *EtcdResolver) Close() {
	r.cancel()
}
