package registry

import "google.golang.org/grpc/resolver"

var exampleScheme = "example"
var exampleServiceName = "resolver.example.grpc.io"

var backendAddr = "localhost:50051"

type etcdv3ResolverBuilder struct{}

func (*etcdv3ResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	r := &etcdv3Resolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			exampleServiceName: {backendAddr},
		},
	}
	r.start()
	return r, nil
}
func (*etcdv3ResolverBuilder) Scheme() string { return exampleScheme }

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
func (*etcdv3Resolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*etcdv3Resolver) Close()                                  {}

func init() {
	resolver.Register(&etcdv3ResolverBuilder{})
}
