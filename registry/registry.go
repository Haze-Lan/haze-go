package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Haze-Lan/haze-go/event"
	"github.com/Haze-Lan/haze-go/option"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"sync"
	"time"
)

var log = grpclog.Component("registry")
var timeOut = time.Duration(3) * time.Second

type Registry interface {
	RegisterService(context.Context, *Instance) error
	UnregisterService(context.Context, string) error
	ListServices(context.Context, string) ([]*Instance, error)
	WatchServices(ctx context.Context, name string) (err error)
}

type etcdv3Registry struct {
	client            *clientv3.Client
	rmu               *sync.RWMutex
	opt               *option.RegistryOptions
	currentServiceKey string
}

func NewRegistry() *etcdv3Registry {
	resolver.Register(&Etcdv3ResolverBuilder{})
	config := clientv3.Config{
		Endpoints:   option.RegistryOptionsInstance.ServerHost,
		DialTimeout: timeOut,
		Context:     context.TODO(),
	}
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	r := &etcdv3Registry{
		client: client,
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
	r.UnregisterService(context.Background(), r.currentServiceKey)
	err := r.client.Close()
	if err != nil {
		log.Error(err)
	}
}

func (r *etcdv3Registry) RegisterService(ctx context.Context, info *Instance) error {
	key := registerKey(info)
	r.currentServiceKey = key
	val := registerValue(info)
	opOptions := make([]clientv3.OpOption, 0)
	_, err := r.client.Put(ctx, key, val, opOptions...)
	if err != nil {
		log.Errorf("register service  fail,%s", err.Error())
		return err
	}
	log.Infof("register service %s", info.Name)
	return nil
}

func (r *etcdv3Registry) UnregisterService(ctx context.Context, key string) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeOut)
		defer cancel()
	}
	_, err := r.client.Delete(ctx, key)
	if err == nil {
		log.Infof("Unloading service %s", key)
		return nil
	}
	return err
}

func (r *etcdv3Registry) ListServices(ctx context.Context, name string) (services []*Instance, err error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeOut)
		defer cancel()
	}
	resp, err := r.client.Get(ctx, name, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	var ss = make([]*Instance, 0)
	for _, ev := range resp.Kvs {
		ins := &Instance{}
		json.Unmarshal(ev.Value, ins)
		ss = append(ss, ins)
	}
	return ss, nil
}
func (r *etcdv3Registry) WatchServices(ctx context.Context, name string) (err error) {
	cacelctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()
	resp, err := r.client.Get(cacelctx, name, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	for _, ev := range resp.Kvs {
		event.GlobalEventBus.Publish(event.EVENT_TOPIC_SERVICE_CHANGE, InstanceEvent{
			eventType: InstanceEventAsAdd,
			kv:        ev,
		})
	}
	var watchChan = r.client.Watch(context.Background(), name, clientv3.WithPrefix(), clientv3.WithRev(resp.Header.Revision+1))
	for {
		select {
		case <-cacelctx.Done():
			return nil
		case resp_ := <-watchChan:
			err = resp_.Err()
			if err != nil {
				return err
			}
			for _, ev := range resp_.Events {
				if ev.IsCreate() {
					event.GlobalEventBus.Publish(event.EVENT_TOPIC_SERVICE_CHANGE, InstanceEvent{
						eventType: InstanceEventAsAdd,
						kv:        ev.Kv,
					})
				} else if ev.IsModify() {
					event.GlobalEventBus.Publish(event.EVENT_TOPIC_SERVICE_CHANGE, InstanceEvent{
						eventType: InstanceEventAsUpdate,
						kv:        ev.Kv,
					})
				} else if ev.Type == mvccpb.DELETE {
					event.GlobalEventBus.Publish(event.EVENT_TOPIC_SERVICE_CHANGE, InstanceEvent{
						eventType: InstanceEventAsDelete,
						kv:        ev.Kv,
					})
				}
			}
		}
	}
	return nil
}

func registerKey(ins *Instance) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", ins.Name, ins.Region, ins.Zone, ins.Namespace, ins.Address)
}

func registerValue(ins *Instance) string {
	val, _ := json.Marshal(ins)
	return string(val)
}
