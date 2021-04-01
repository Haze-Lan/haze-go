package discovery

//微服务实体类
type Service struct {
	//服务名称
	name string
	//服务唯一标识
	id string
	//监控检查地址
	check string
	//元数据  作为服务选择依据
	meta map[string]string
	//标签
	label []string
}
