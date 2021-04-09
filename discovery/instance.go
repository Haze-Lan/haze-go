package discovery

import (
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
)

type Instance struct {
	Weight float64
	Ip     string
	Port   uint64
	//服务名称
	Name string
	//服务唯一标识
	Id string
	//监控检查地址
	Check string
	//元数据  作为服务选择依据
	Meta map[string]string
}

func NewService(opt *option.ServerOptions) *Instance {
	discoveryOptions, _ := option.LoadDiscoveryOptions()
	ip := utils.GetIp()
	var defaultService = &Instance{Meta: make(map[string]string, 0)}
	defaultService = defaultService.WithID(discoveryOptions.InstanceId).WithPort(opt.Port).WithIp(ip).WithWeight(discoveryOptions.InstanceWeight).WithName(opt.Name).WithIp(opt.Host)
	return defaultService
}
func (s *Instance) WithWeight(w float64) *Instance {
	s.Weight = w
	return s
}
func (s *Instance) WithIp(ip string) *Instance {
	s.Ip = ip
	return s
}
func (s *Instance) WithID(id string) *Instance {
	if len(id) < 1 {
		idItem, _ := uuid.NewV4()
		id = idItem.String()
	}
	s.Id = id
	return s
}
func (s *Instance) WithPort(port uint64) *Instance {
	s.Port = port
	return s
}
func (s *Instance) WithName(name string) *Instance {
	s.Name = name
	return s
}
func (s *Instance) WithCheck(check string) *Instance {
	s.Check = check
	return s
}
func (s *Instance) WithMeta(key string, value string) *Instance {
	s.Meta[key] = value
	return s
}
