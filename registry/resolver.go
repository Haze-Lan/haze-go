package registry

import (
	"encoding/json"
	"github.com/Haze-Lan/haze-go/event"
	"google.golang.org/grpc/resolver"
)

type Etcdv3ResolverBuilder struct {
}

func (*Etcdv3ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &etcdv3Resolver{
		target:     target,
		cc:         cc,
		addrsStore: make(map[string]map[string]string),
	}
	r.start()
	return r, nil
}
func (*Etcdv3ResolverBuilder) Scheme() string { return "etcd" }

type etcdv3Resolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string]map[string]string
}

func (r *etcdv3Resolver) start() {
	event.GlobalEventBus.Subscribe(event.EVENT_TOPIC_SERVICE_CHANGE, func(data interface{}) {
		eventData := data.(InstanceEvent)
		switch eventData.eventType {
		case InstanceEventAsDelete:
			ins := &Instance{}
			json.Unmarshal(eventData.kv.Value, ins)
			addrStrs := r.addrsStore[r.target.Endpoint]
			if addrStrs != nil {
				delete(addrStrs, ins.Id)
			}
			addrs := make([]resolver.Address, 0)
			for _, v := range addrStrs {
				addrs = append(addrs, resolver.Address{Addr: v})
			}
			r.cc.UpdateState(resolver.State{Addresses: addrs})

		case InstanceEventAsAdd:
			ins := &Instance{}
			json.Unmarshal(eventData.kv.Value, ins)
			addrStrs := r.addrsStore[r.target.Endpoint]
			if addrStrs == nil {
				addrStrs = make(map[string]string, 0)
				r.addrsStore[r.target.Endpoint] = addrStrs
			}
			addrStrs[ins.Id] = ins.Address
			addrs := make([]resolver.Address, 0)
			for _, v := range addrStrs {
				addrs = append(addrs, resolver.Address{Addr: v})
			}
			r.cc.UpdateState(resolver.State{Addresses: addrs})
		case InstanceEventAsUpdate:
			ins := &Instance{}
			json.Unmarshal(eventData.kv.Value, ins)
			addrStrs := r.addrsStore[r.target.Endpoint]
			if addrStrs == nil {
				addrStrs = make(map[string]string, 0)
				r.addrsStore[r.target.Endpoint] = addrStrs
			}
			addrStrs[ins.Id] = ins.Address
			addrs := make([]resolver.Address, 0)
			for _, v := range addrStrs {
				addrs = append(addrs, resolver.Address{Addr: v})
			}
			r.cc.UpdateState(resolver.State{Addresses: addrs})
		}
	})

}
func (*etcdv3Resolver) ResolveNow(o resolver.ResolveNowOptions) {
	log.Info("Analytical services")
}
func (*etcdv3Resolver) Close() {}
