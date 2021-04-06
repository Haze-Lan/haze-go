package discovery

type Discovery interface {
	//初始化
	Init(opt *DiscoveryOption) error
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

func NewDiscovery(opt *DiscoveryOption) Discovery {
	var dis Discovery
	switch opt.DiscoveryType {
	case "nacos":
		dis = &Nacos{}
	}
	err:=dis.Init(opt)
	if err!=nil {
		log.Fatalln(err)
	}
	return dis
}
