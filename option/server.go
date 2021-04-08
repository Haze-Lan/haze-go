package option

//服务配置
type ServerOptions struct {
	Host string `properties:"server.host,default=127.0.0.1"`
	Port uint64 `properties:"server.port,default=80"`
	Name string `properties:"server.name,default=haze"`
}
