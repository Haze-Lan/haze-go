package discovery

import (
	"github.com/Haze-Lan/haze-go/logger"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var log = logger.LoggerFactory("nacos")

type Nacos struct {
	client naming_client.INamingClient
}

func (dis *Nacos) Init(opt *DiscoveryOption) error {
	log.Infof("加载服务  %s", "nacos")
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(""), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("\\nacos\\log"),
		constant.WithCacheDir("\\nacos\\cache"),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel("debug"),
	)
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			"127.0.0.1",
			8848,
			constant.WithScheme("http"),
			constant.WithContextPath("/nacos"),
		),
	}
	var err error
	dis.client, err = clients.NewNamingClient(vo.NacosClientParam{ClientConfig: &clientConfig, ServerConfigs: serverConfigs})
	if err != nil {
		return err
	}
	return nil
}

func (dis *Nacos) Destroy() error {

	return nil
}

func (dis *Nacos) RegisterInstance(service *Service) error {
	success, err := dis.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        80,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc":"shanghai"},
		ClusterName: "DEFAULT", // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP",   // 默认值DEFAULT_GROUP
	})
	if err!=nil {
		return err
	}
	if success {

	}
	return nil
}

func (dis *Nacos) DeregisterInstance(service *Service) error {
	return nil
}
func (dis *Nacos) GetService() (services []Service, err error) {
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
