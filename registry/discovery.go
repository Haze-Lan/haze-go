package registry

import (
	"context"
	"github.com/Haze-Lan/haze-go/logger"
	"github.com/Haze-Lan/haze-go/option"
	"go.etcd.io/etcd/clientv3"
	"time"
)

type Discovery interface {

	//销毁
	Destroy(ctx context.Context) error

	RegisterInstance(ctx context.Context, service *Instance) error

	DeregisterInstance(ctx context.Context, service *Instance) error

	GetService(ctx context.Context) (services []Instance, err error)

	SelectInstances(ctx context.Context) (instances []Instance, err error)

	//监听服务
	SubscribeService(ctx context.Context) error
	//取消监听
	UnsubscribeService(ctx context.Context) error
}

var discoveryOptions *option.DiscoveryOptions
var log = logger.Factory("registry")

func NewDiscovery() Discovery {
	discoveryOptions, err := option.LoadDiscoveryOptions()
	if err != nil {
		log.Fatal(err.Error())
	}
	config := clientv3.Config{
		Endpoints:   discoveryOptions.ServerHost,
		DialTimeout: 10 * time.Second,
		Context:     context.TODO(),
	}
	client, err := clientv3.New(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Info("Discovery Initialization completed")
	return &Etcd{opt: discoveryOptions, client: client, ctx: context.Background()}
}
