package option

import (
	"flag"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/magiconair/properties"
)

var prop *properties.Properties

func init() {
	path, _ := utils.GetCurrentDirectory()
	prop = properties.MustLoadFile(path+"/application.properties", properties.UTF8)
	prop.MustFlag(flag.CommandLine)
}

func LoadServerOptions() (opt *ServerOptions, err error) {
	var serverOptionsInstance = &ServerOptions{}
	err = prop.Decode(serverOptionsInstance)
	if err != nil {
		return nil, err
	}
	return serverOptionsInstance, nil
}
func LoadDiscoveryOptions() (opt *DiscoveryOptions, err error) {
	var discoveryOptionsInstance = &DiscoveryOptions{}
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
