package discovery

type Discovery interface {

	//销毁
	Destroy() error

	RegisterInstance(service *Service) error

	DeregisterInstance(service *Service) error

	GetService() (services []Service, err error)

	SelectInstances() (instances []Instance, err error)

	//监听服务
	SubscribeService() error
	//取消监听
	UnsubscribeService() error
}

func NewDiscovery(disType string, opts... DiscoveryOption) Discovery {
	for _,opt := range opts {
		opt.apply(defaultDiscoveryOption)
	}
	var dis Discovery
	switch disType {
	case "nacos":
		newNacos(defaultDiscoveryOption)
	}
	return dis
}
