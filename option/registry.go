package option

// RegistryOptions 注册中心配置
type RegistryOptions struct {
	ServerHost           []string          `properties:"discovery.server.host,default=127.0.0.1:2379"`
	ServerNameSpace      string            `properties:"discovery.server.namespace,default=haodoings.com"`
	ServerRegion         string            `properties:"discovery.server.region,default=default"`
	ServerZone           string            `properties:"discovery.server.zone,default=default"`
	InstanceIntervalTime uint64            `properties:"discovery.server.instance.interval.time,default=1"`
	InstanceName         string            `properties:"discovery.server.instance.name,default=haze"`
	InstanceWeight       float64           `properties:"discovery.server.instance.weight,default=10"`
	InstanceMetadata     map[string]string `properties:"discovery.server.instance.metadata,default="`
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
