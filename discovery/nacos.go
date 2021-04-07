package discovery

import (
	"github.com/Haze-Lan/haze-go/logger"
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
}

func newNacos(opts *discoveryOption) Discovery {
	var nacos = &Nacos{}
	path, err := os.Executable()
	if err != nil {
		log.Error(err)
	}
	dir := filepath.Dir(path)
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(utils.ValueOfString(opts.meta["namespace"], "")),
		constant.WithTimeoutMs(utils.ValueOfInt(opts.meta["timeout"], 5000)),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(utils.ValueOfString(opts.meta["logDir"], dir+"\\nacos\\log")),
		constant.WithCacheDir(utils.ValueOfString(opts.meta["cacheDir"], dir+"\\nacos\\cache")),
		constant.WithRotateTime("1h"),
		constant.WithMaxAge(3),
		constant.WithLogLevel(utils.ValueOfString(opts.meta["logLevel"], "info")),
	)
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			utils.ValueOfString(opts.serverHost, "127.0.0.1"),
			utils.ValueOfInt(opts.serverPort, 8848),
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

func (dis *Nacos) RegisterInstance(service *Service) error {
	success, err := dis.client.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          "10.0.0.11",
		Port:        80,
		ServiceName: "demo.go",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	if err != nil {
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
