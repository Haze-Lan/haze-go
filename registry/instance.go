package registry

import (
	"fmt"
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/utils"
	"go.etcd.io/etcd/api/v3/mvccpb"
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

func NewInstance(addr string) *Instance {
	var id = utils.Hash(fmt.Sprintf("%s-%s-%s-%s-%s", option.RegistryOptionsInstance.ServerRegion, option.RegistryOptionsInstance.ServerZone, option.RegistryOptionsInstance.ServerNameSpace, option.RegistryOptionsInstance.InstanceName, addr))
	si := &Instance{
		Name:      option.RegistryOptionsInstance.InstanceName,
		Id:        id,
		Weight:    option.RegistryOptionsInstance.InstanceWeight,
		Address:   addr,
		Enable:    true,
		Metadata:  make(map[string]string),
		Region:    option.RegistryOptionsInstance.ServerRegion,
		Zone:      option.RegistryOptionsInstance.ServerZone,
		Namespace: option.RegistryOptionsInstance.ServerNameSpace,
	}
	si.Metadata["time"] = time.Now().String()
	si.Metadata["version"] = "1.0"
	return si
}

type InstanceEventType int

const InstanceEventAsAdd = InstanceEventType(1)
const InstanceEventAsDelete = InstanceEventType(2)
const InstanceEventAsUpdate = InstanceEventType(3)

type InstanceEvent struct {
	eventType InstanceEventType
	kv        *mvccpb.KeyValue
}
