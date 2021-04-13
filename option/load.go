package option

import (
	"flag"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/magiconair/properties"
	"google.golang.org/grpc/grpclog"
)

var log = grpclog.Component("option")
var prop *properties.Properties
var serverOptionsInstance *ServerOptions

func init() {
	path, _ := utils.GetCurrentDirectory()
	prop = properties.MustLoadFile(path+"/application.properties", properties.UTF8)
	prop.MustFlag(flag.CommandLine)
}

func LoadServerOptions() (opt *ServerOptions, err error) {
	serverOptionsInstance = &ServerOptions{}
	err = prop.Decode(serverOptionsInstance)
	if err != nil {
		return nil, err
	}
	return serverOptionsInstance, nil
}
func LoadDiscoveryOptions() (opt *RegistryOptions, err error) {
	var discoveryOptionsInstance = &RegistryOptions{}
	err = prop.Decode(discoveryOptionsInstance)
	if err != nil {
		return nil, err
	}

	return discoveryOptionsInstance, nil
}

func LoadLoggingOptions() (opt *LoggingOptions, err error) {
	var loggingOptionsInstance = &LoggingOptions{}
	err = prop.Decode(loggingOptionsInstance)
	if err != nil {
		return nil, err
	}
	return loggingOptionsInstance, nil
}
