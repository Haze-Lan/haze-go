package discovery



//ETCD
type Etcd struct {


}

//初始化 配置 并链接注册中心 完成就绪
func (etcd *Etcd) Init() error{
	return nil
}
//注册服务
func (etcd *Etcd) Install(service... Service) error {
	return nil
}
//卸载服务
func (etcd *Etcd) Uninstall(service Service) error{
	return nil;
}
//查询服务
func (etcd *Etcd) Instances(name string)(services []Service,err error){

	return nil,nil
}
//组件注销
func (etcd *Etcd) Destroy(service Service) error {
	return nil
}