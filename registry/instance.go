package registry

import (
	"context"
	"fmt"
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
)



type Option func(c *Instance)

// ServiceConfigurator represents service configurator
type ConfigInfo struct {
	Routes []Route
}

// ServiceInfo represents service info
type Instance struct {
	Name     string               `json:"name"`
	AppID    string               `json:"appId"`
	Scheme   string               `json:"scheme"`
	Address  string               `json:"address"`
	Weight   float64              `json:"weight"`
	Enable   bool                 `json:"enable"`
	Healthy  bool                 `json:"healthy"`
	Metadata map[string]string    `json:"metadata"`
	Region   string               `json:"region"`
	Zone     string               `json:"zone"`
	Kind     ServiceKind `json:"kind"`
	// Deployment 部署组: 不同组的流量隔离
	// 比如某些服务给内部调用和第三方调用，可以配置不同的deployment,进行流量隔离
	Deployment string `json:"deployment"`
	// Group 流量组: 流量在Group之间进行负载均衡
	Group    string              `json:"group"`
	Services map[string]*Service `json:"services" toml:"services"`
}

// Service ...
type Service struct {
	Namespace string            `json:"namespace" toml:"namespace"`
	Name      string            `json:"name" toml:"name"`
	Labels    map[string]string `json:"labels" toml:"labels"`
	Methods   []string          `json:"methods" toml:"methods"`
}

// Label ...
func (si Instance) Label() string {
	return fmt.Sprintf("%s://%s", si.Scheme, si.Address)
}

// Server ...
type Server interface {
	Serve() error
	Stop() error
	GracefulStop(ctx context.Context) error
	Info() *Instance
}



func ApplyOptions(options ...Option) Instance {
	info := defaultServiceInfo()
	for _, option := range options {
		option(&info)
	}
	return info
}

func WithMetaData(key, value string) Option {
	return func(c *Instance) {
		c.Metadata[key] = value
	}
}

func WithScheme(scheme string) Option {
	return func(c *Instance) {
		c.Scheme = scheme
	}
}

func WithAddress(address string) Option {
	return func(c *Instance) {
		c.Address = address
	}
}

func WithKind(kind ServiceKind) Option {
	return func(c *Instance) {
		c.Kind = kind
	}
}

func defaultServiceInfo() Instance {
	si := Instance{
		Name:       pkg.Name(),
		AppID:      pkg.AppID(),
		Weight:     100,
		Enable:     true,
		Healthy:    true,
		Metadata:   make(map[string]string),
		Region:     pkg.AppRegion(),
		Zone:       pkg.AppZone(),
		Kind:       0,
		Deployment: "",
		Group:      "",
	}
	si.Metadata["appMode"] = pkg.AppMode()
	si.Metadata["appHost"] = pkg.AppHost()
	si.Metadata["startTime"] = pkg.StartTime()
	si.Metadata["buildTime"] = pkg.BuildTime()
	si.Metadata["appVersion"] = pkg.AppVersion()
	si.Metadata["jupiterVersion"] = pkg.JupiterVersion()
	return si
}
© 2021 GitHub, Inc.