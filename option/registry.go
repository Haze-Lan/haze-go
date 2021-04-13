package option

//注册中心配置
type RegistryOptions struct {
	InstanceIntervalTime uint64            `properties:"registry.server.instance.interval.time,default=1"`
	ServerHost           []string          `properties:"registry.server.host,default=127.0.0.1:2379"`
	ServerNameSpace      string            `properties:"registry.server.namespace,default=haodoings.com"`
	InstanceName         string            `properties:"registry.server.instance.name,default=haze"`
	InstanceRegion       string            `properties:"registry.server.instance.region,default=default"`
	InstanceZone         string            `properties:"registry.server.instance.zone,default=default"`
	InstanceId           string            `properties:"registry.server.instance.id,default=1"`
	InstanceWeight       float64           `properties:"registry.server.instance.weight,default=10"`
	InstanceAddress      string            `properties:"registry.server.instance.address,default=127.0.0.1"`
	InstanceHealth       string            `properties:"registry.server.instance.health,default="`
	InstanceMetadata     map[string]string `properties:"registry.server.instance.metadata,default="`
}

type RegistryOptionsFun interface {
	Apply(*RegistryOptions)
}
type funcRegistryOptionsFun struct {
	f func(*RegistryOptions)
}

func (fdo *funcRegistryOptionsFun) Apply(do *RegistryOptions) {
	fdo.f(do)
}

func newFuncRegistryOptionsFun(f func(*RegistryOptions)) *funcRegistryOptionsFun {
	return &funcRegistryOptionsFun{
		f: f,
	}
}
func WithId(s string) RegistryOptionsFun {
	return newFuncRegistryOptionsFun(func(o *RegistryOptions) {
		o.InstanceId = s
	})
}
func WithInstanceAddress(s string) RegistryOptionsFun {
	return newFuncRegistryOptionsFun(func(o *RegistryOptions) {
		o.InstanceAddress = s
	})
}
func WithInstanceName(s string) RegistryOptionsFun {
	return newFuncRegistryOptionsFun(func(o *RegistryOptions) {
		o.InstanceName = s
	})
}
