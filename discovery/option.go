package discovery

import (
	"flag"
	"github.com/Haze-Lan/haze-go/utils"
	"strconv"
	"strings"
)

var localIp4 = utils.GetIp()
var defaultDiscoveryOption = &discoveryOption{
	serverHost: localIp4,
	serverPort: 80,
	meta:       make(map[string]string),
}

type discoveryOption struct {
	serverHost string
	serverPort uint64
	meta       map[string]string
}


var unParseOpts []DiscoveryOption

func init() {
	unParseOpts = make([]DiscoveryOption, 0, 10)
	flag.CommandLine.Func("discovery.sever.host", "注册中心地址", func(s string) error {
		unParseOpts = append(unParseOpts, WithServerHost(s))
		return nil
	})
	flag.CommandLine.Func("discovery.sever.post", "注册中心端口", func(s string) error {
		unParseOpts = append(unParseOpts, WithServerPort(s))
		return nil
	})
	flag.Parse()
}




type DiscoveryOption interface {
	apply(*discoveryOption)
}

type funcDiscoveryOption struct {
	f func(*discoveryOption)
}

func (fdo *funcDiscoveryOption) apply(do *discoveryOption) {
	fdo.f(do)
}
func newFuncDiscoveryOption(f func(*discoveryOption)) *funcDiscoveryOption {
	return &funcDiscoveryOption{
		f: f,
	}
}

func WithServerHost(host string) DiscoveryOption {
	return newFuncDiscoveryOption(func(o *discoveryOption) {
		if len(strings.TrimSpace(host)) > 0 {
			o.serverHost = host
		}
	})
}
func WithServerPort(port string) DiscoveryOption {
	return newFuncDiscoveryOption(func(o *discoveryOption) {
		o.serverPort,_ =  strconv.ParseUint(port,10,64)
	})
}
