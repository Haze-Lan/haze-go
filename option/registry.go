package option

// RegistryOptions 注册中心配置
type RegistryOptions struct {
	ServerHost           []string          `properties:"registry.server.host,default=127.0.0.1:2379"`
	ServerNameSpace      string            `properties:"registry.server.namespace,default=haodoings.com"`
	ServerRegion         string            `properties:"registry.server.region,default=default"`
	ServerZone           string            `properties:"registry.server.zone,default=default"`
	InstanceIntervalTime uint64            `properties:"registry.server.instance.interval.time,default=1"`
	InstanceName         string            `properties:"registry.server.instance.name,default=haze"`
	InstanceWeight       float64           `properties:"registry.server.instance.weight,default=10"`
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
func WithInstanceName(s string) RegistryOptionsFun {
	return newFuncRegistryOptionsFun(func(o *RegistryOptions) {
		o.InstanceName = s
	})
}
