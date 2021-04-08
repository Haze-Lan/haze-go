package discovery

import (
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
)

type Instance struct {
	weight float64
	ip     string
	port   uint64
	//服务名称
	name string
	//服务唯一标识
	id string
	//监控检查地址
	check string
	//元数据  作为服务选择依据
	meta map[string]string
}

func NewService(opt *option.ServerOptions) *Instance {
	discoveryOptions, _ := option.LoadDiscoveryOptions()
	ip := utils.GetIp()
	var defaultService = &Instance{meta: make(map[string]string, 0)}
	defaultService = defaultService.WithID(discoveryOptions.InstanceId).WithPort(opt.Port).WithIp(ip).WithWeight(discoveryOptions.InstanceWeight).WithName(opt.Name).WithIp(opt.Host)
	return defaultService
}
func (s *Instance) WithWeight(w float64) *Instance {
	s.weight = w
	return s
}
func (s *Instance) WithIp(ip string) *Instance {
	s.ip = ip
	return s
}
func (s *Instance) WithID(id string) *Instance {
	if len(id) < 1 {
		idItem, _ := uuid.NewV4()
		id = idItem.String()
	}
	s.id = id
	return s
}
func (s *Instance) WithPort(port uint64) *Instance {
	s.port = port
	return s
}
func (s *Instance) WithName(name string) *Instance {
	s.name = name
	return s
}
func (s *Instance) WithCheck(check string) *Instance {
	s.check = check
	return s
}
func (s *Instance) WithMeta(key string, value string) *Instance {
	s.meta[key] = value
	return s
}
