package option

//注册中心配置
type DiscoveryOptions struct {


	//心跳间隔 秒
	InstanceIntervalTime uint64 `properties:"discovery.server.instance.interval.time,default=1"`
	ServerHost       []string          `properties:"registry.server.host,default=127.0.0.1:2379"`
	InstanceId       string            `properties:"registry.server.instance.id,default=1"`
	InstanceWeight   float64           `properties:"registry.server.instance.weight,default=10"`
	InstanceIp       string            `properties:"registry.server.instance.ip,default=127.0.0.1"`
	InstancePort     uint64            `properties:"registry.server.instance.port,default=80"`
	InstanceHealth   string            `properties:"registry.server.instance.health,default="`
	InstanceMetadata map[string]string `properties:"registry.server.instance.metadata,default="`

}
