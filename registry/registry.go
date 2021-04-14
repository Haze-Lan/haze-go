package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Haze-Lan/haze-go/event"
	"github.com/Haze-Lan/haze-go/option"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)


var log = grpclog.Component("registry")

type Registry interface {
	RegisterService(context.Context, *Instance) error
	UnregisterService(context.Context, *Instance) error
	ListServices(context.Context, string, string) ([]*Instance, error)
}
type etcdv3Registry struct {
	client *clientv3.Client
	kvs    sync.Map
	cancel context.CancelFunc
	rmu    *sync.RWMutex
	opt    *option.RegistryOptions
}



func NewRegistry() *etcdv3Registry {
	resolver.Register(&Etcdv3ResolverBuilder{})
	config := clientv3.Config{
		Endpoints:   option.RegistryOptionsInstance.ServerHost,
		DialTimeout: 10 * time.Second,
		Context:     context.TODO(),
	}
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	r := &etcdv3Registry{
		client: client,
		kvs:    sync.Map{},
		rmu:    &sync.RWMutex{},
		opt:    option.RegistryOptionsInstance,
	}
	log.Info("discovery initialization completed")
	event.GlobalEventBus.Subscribe(event.EVENT_TOPIC_SERVER_QUIT, func(data interface{}) {
		r.Stop()
	})
	return r
}

func (r *etcdv3Registry) Stop() {
	log.Infof("close the %s component", "registry")
	 err:=r.client.Close()
	if err!=nil {
		log.Error(err)
	}
}

func (r *etcdv3Registry) RegisterService(ctx context.Context, info *Instance) error {
	key := r.registerKey(info)
	val := r.registerValue(info)
	opOptions := make([]clientv3.OpOption, 0)
	_, err := r.client.Put(ctx, key, val, opOptions...)
	if err != nil {
		log.Errorf("register service %s", err.Error())
		return err
	}
	log.Infof("register service %s %s", key, val)
	r.kvs.Store(key, val)
	return nil
}

func (r *etcdv3Registry) UnregisterService(ctx context.Context, info *Instance) error {
	return r.unregister(ctx, r.registerKey(info))
}

func (r *etcdv3Registry) ListServices(ctx context.Context, name string, scheme string) (services []*Instance, err error) {
	return nil, nil
}
func (r *etcdv3Registry) unregister(ctx context.Context, key string) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10)
		defer cancel()
	}
	_, err := r.client.Delete(ctx, key)
	if err == nil {
		r.kvs.Delete(key)
	}
	return err
}

func (r *etcdv3Registry) registerKey(ins *Instance) string {
	return fmt.Sprintf("/%s/%s/%s/%s://%s", ins.Region, ins.Zone, ins.Namespace, ins.Name, ins.Address)
}

func (r *etcdv3Registry) registerValue(ins *Instance) string {
	val, _ := json.Marshal(ins)
	return string(val)
}
