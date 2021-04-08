package discovery

import (
	"github.com/Haze-Lan/haze-go/logger"
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"os"
	"path/filepath"
)

var log = logger.LoggerFactory("nacos")

type Nacos struct {
	client naming_client.INamingClient
	opt    *option.DiscoveryOptions
}

func newNacos(opts *option.DiscoveryOptions) Discovery {
	var nacos = &Nacos{}
	path, err := os.Executable()
	if err != nil {
		log.Error(err)
	}
	dir := filepath.Dir(path)
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(utils.ValueOfString(opts.InstanceMetadata["namespace"], "")),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(utils.ValueOfString(opts.InstanceMetadata["logDir"], dir+"\\nacos\\log")),
		constant.WithCacheDir(utils.ValueOfString(opts.InstanceMetadata["cacheDir"], dir+"\\nacos\\cache")),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel(utils.ValueOfString(opts.InstanceMetadata["logLevel"], "info")),
	)
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			utils.ValueOfString(opts.ServerHost, "127.0.0.1"),
			utils.ValueOfInt(opts.ServerPort, 8848),
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
	nacos.client, err = clients.NewNamingClient(vo.NacosClientParam{ClientConfig: &clientConfig, ServerConfigs: serverConfigs})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return nacos
}

func (dis *Nacos) Destroy() error {

	return nil
}

func (dis *Nacos) RegisterInstance(service *Instance) error {
	_, err := dis.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          service.ip,
		Port:        service.port,
		ServiceName: service.name,
		Weight:      service.weight,
		Enable:      true,
		Healthy:     len(service.check) != 0,
		Ephemeral:   true,
		Metadata:    service.meta,
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	return nil
}

func (dis *Nacos) DeregisterInstance(service *Instance) error {
	_, err := dis.client.DeregisterInstance(vo.DeregisterInstanceParam{
		Ip:          service.ip,
		Port:        service.port,
		ServiceName: service.name,
		Ephemeral:   true,
		Cluster:     "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	return nil
}
func (dis *Nacos) GetService() (services []Instance, err error) {
	return nil, nil
}
func (dis *Nacos) SelectInstances() (instances []Instance, err error) {
	return nil, nil
}

func (dis *Nacos) SubscribeService() error {
	return nil
}
func (dis *Nacos) UnsubscribeService() error {
	return nil
}
