package registry

import (
	"google.golang.org/grpc/resolver"
)



type Etcdv3ResolverBuilder struct{

}

func (*Etcdv3ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	r := &etcdv3Resolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{"account-haze":[]string{"127.0.0.1:80"}},
	}
	r.start()
	return r, nil
}
func (*Etcdv3ResolverBuilder) Scheme() string { return "etcd" }

type etcdv3Resolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *etcdv3Resolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*etcdv3Resolver) ResolveNow(o resolver.ResolveNowOptions) {
	log.Info(o)
}
func (*etcdv3Resolver) Close()                                  {}


