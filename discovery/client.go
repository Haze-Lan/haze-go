package discovery

//服务发现接口
type Client interface {
	//初始化 配置 并链接注册中心 完成就绪
	Init() error
	//注册服务
	Install(service ...Service) error
	//卸载服务
	Uninstall(service Service) error
	//查询服务
	Instances(name string) (services []Service, err error)
	//组件注销
	Destroy(service Service) error
}
