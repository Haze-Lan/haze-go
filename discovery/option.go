package discovery

import (
	"github.com/Haze-Lan/haze-go/utils"
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
func WithServerPort(port uint64) DiscoveryOption {
	return newFuncDiscoveryOption(func(o *discoveryOption) {
		o.serverPort = port
	})
}
