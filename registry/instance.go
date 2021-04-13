package registry

import (
	"github.com/Haze-Lan/haze-go/option"
	"time"
)

type Instance struct {
	Namespace string            `json:"namespace" `
	Name      string            `json:"name"`
	Id        string            `json:"id"`
	Address   string            `json:"address"`
	Weight    float64           `json:"weight"`
	Enable    bool              `json:"enable"`
	Metadata  map[string]string `json:"metadata"`
	Region    string            `json:"region"`
	Zone      string            `json:"zone"`
	Methods   []string          `json:"methods" `
}

func NewInstance(opts ...option.RegistryOptionsFun) *Instance {
	for _, opt := range opts {
		opt.Apply(options)
	}
	si := &Instance{
		Name:     options.InstanceName,
		Id:       options.InstanceId,
		Weight:   options.InstanceWeight,
		Enable:   true,
		Metadata: make(map[string]string),
		Region:   options.InstanceRegion,
		Zone:     options.InstanceZone,
	}
	si.Metadata["time"] = time.Now().String()
	si.Metadata["version"] = "1.0"
	return si
}
