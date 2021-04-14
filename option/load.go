package option

import (
	"flag"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/magiconair/properties"
	"google.golang.org/grpc/grpclog"
)

var log = grpclog.Component("option")
var prop *properties.Properties
var ServerOptionsInstance *ServerOptions
var RegistryOptionsInstance *RegistryOptions
var LoggingOptionsInstance *LoggingOptions

func init() {
	path, _ := utils.GetCurrentDirectory()
	filePath:=flag.CommandLine.String("c", path + "/application.properties", "configuration file path")
	flag.Parse()
	prop = properties.MustLoadFile(*filePath, properties.UTF8)
	prop.MustFlag(flag.CommandLine)
	LoggingOptionsInstance = &LoggingOptions{}
	err := prop.Decode(LoggingOptionsInstance)
	if err != nil {
		log.Fatal(err)
	}
	ServerOptionsInstance = &ServerOptions{}
	err = prop.Decode(ServerOptionsInstance)
	if err != nil {
		log.Fatal(err)
	}
	RegistryOptionsInstance = &RegistryOptions{}
	err = prop.Decode(RegistryOptionsInstance)
	if err != nil {
		log.Fatal(err)
	}
}
