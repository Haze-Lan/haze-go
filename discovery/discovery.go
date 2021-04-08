package discovery

import "github.com/Haze-Lan/haze-go/option"

type Discovery interface {

	//销毁
	Destroy() error

	RegisterInstance(service *Instance) error

	DeregisterInstance(service *Instance) error

	GetService() (services []Instance, err error)

	SelectInstances() (instances []Instance, err error)

	//监听服务
	SubscribeService() error
	//取消监听
	UnsubscribeService() error
}

var discoveryOptions *option.DiscoveryOptions

func NewDiscovery() Discovery {
	discoveryOptions, err := option.LoadDiscoveryOptions()
	if err != nil {
		log.Fatalln(err.Error())
	}
	var dis Discovery
	switch discoveryOptions.ServerType {
	case "nacos":
		dis = newNacos(discoveryOptions)
	}
	return dis
}
