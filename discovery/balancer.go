package discovery

import (
	"fmt"
	"github.com/go-errors/errors"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/connectivity"
)

const randomBalancerName = "random"

func newRandomBalancerBuilder() balancer.Builder {
	return &RandomBuilder{}
}

type RandomBuilder struct{}

func (*RandomBuilder) Build(cc balancer.ClientConn, opt balancer.BuildOptions) balancer.Balancer {
	return &RandomBalancer{cc: cc}
}

func (*RandomBuilder) Name() string {
	return randomBalancerName
}

type RandomBalancer struct {
	state connectivity.State
	cc    balancer.ClientConn
	sc    balancer.SubConn
}

func (b *RandomBalancer) ResolverError(err error) {
	switch b.state {
	case connectivity.TransientFailure, connectivity.Idle, connectivity.Connecting:
		b.cc.UpdateState(balancer.State{ConnectivityState: connectivity.TransientFailure,
			Picker: &picker{err: fmt.Errorf("name resolver error: %v", err)},
		})
	}
	log.Errorf("RandomBalancer: ResolverError called with error %v", err)
}

func (b *RandomBalancer) UpdateClientConnState(cs balancer.ClientConnState) error {
	if len(cs.ResolverState.Addresses) == 0 {
		b.ResolverError(errors.New("produced zero addresses"))
		return balancer.ErrBadResolverState
	}
	if b.sc == nil {
		var err error
		b.sc, err = b.cc.NewSubConn(cs.ResolverState.Addresses, balancer.NewSubConnOptions{})
		if err != nil {
			if log.V(2) {
				log.Errorf("RandomBalancer: failed to NewSubConn: %v", err)
			}
			b.state = connectivity.TransientFailure
			b.cc.UpdateState(balancer.State{ConnectivityState: connectivity.TransientFailure,
				Picker: &picker{err: fmt.Errorf("error creating connection: %v", err)},
			})
			return balancer.ErrBadResolverState
		}
		b.state = connectivity.Idle
		b.cc.UpdateState(balancer.State{ConnectivityState: connectivity.Idle, Picker: &picker{result: balancer.PickResult{SubConn: b.sc}}})
		b.sc.Connect()
	} else {
		b.cc.UpdateAddresses(b.sc, cs.ResolverState.Addresses)
		b.sc.Connect()
	}
	return nil
}

func (b *RandomBalancer) UpdateSubConnState(sc balancer.SubConn, s balancer.SubConnState) {

	log.Infof("RandomBalancer: UpdateSubConnState: %p, %v", sc, s)

	if b.sc != sc {

		log.Infof("RandomBalancer: ignored state change because sc is not recognized")

		return
	}
	b.state = s.ConnectivityState
	if s.ConnectivityState == connectivity.Shutdown {
		b.sc = nil
		return
	}

	switch s.ConnectivityState {
	case connectivity.Ready, connectivity.Idle:
		b.cc.UpdateState(balancer.State{ConnectivityState: s.ConnectivityState, Picker: &picker{result: balancer.PickResult{SubConn: sc}}})
	case connectivity.Connecting:
		b.cc.UpdateState(balancer.State{ConnectivityState: s.ConnectivityState, Picker: &picker{err: balancer.ErrNoSubConnAvailable}})
	case connectivity.TransientFailure:
		b.cc.UpdateState(balancer.State{
			ConnectivityState: s.ConnectivityState,
			Picker:            &picker{err: s.ConnectionError},
		})
	}
}

func (b *RandomBalancer) Close() {
}

type picker struct {
	result balancer.PickResult
	err    error
}

func (p *picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	return p.result, p.err
}

func init() {
	//balancer.Register(newRandomBalancerBuilder())
}
