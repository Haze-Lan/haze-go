package option

//注册中心配置
type DiscoveryOptions struct {
	ServerHost       string            `properties:"discovery.server.host,default=127.0.0.1"`
	ServerPort       uint64            `properties:"discovery.server.port,default=8848"`
	InstanceId       string            `properties:"discovery.server.instance.id,default=1"`
	InstanceWeight   float64           `properties:"discovery.server.instance.weight,default=10"`
	InstanceIp       string            `properties:"discovery.server.instance.ip,default=127.0.0.1"`
	InstancePort     uint64            `properties:"discovery.server.instance.port,default=80"`
	InstanceHealth   string            `properties:"discovery.server.instance.health,default="`
	InstanceMetadata map[string]string `properties:"discovery.server.instance.metadata,default="`
	ServerType       string            `properties:"discovery.server.type,default=nacos"`
}
