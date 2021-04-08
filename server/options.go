package server

import (
	"flag"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/utils"
	"strconv"
)

var defaultServerOptions *serverOptions
var unParseOpts []ServerOption

func init() {
	defaultServerOptions = &serverOptions{name: utils.GetUUID(), port: 80}
	unParseOpts = make([]ServerOption, 0, 10)
	flag.CommandLine.Func("server.port", "应用启动端口", func(s string) error {
		unParseOpts = append(unParseOpts, WithPort(s))
		return nil
	})
	flag.CommandLine.Func("server.name", "应用名称", func(s string) error {
		unParseOpts = append(unParseOpts, WithName(s))
		return nil
	})
	flag.Parse()
}

type serverOptions struct {
	name string
	port uint64
}
type ServerOption interface {
	apply(*serverOptions)
}
type funcServerOption struct {
	f func(*serverOptions)
}

func (fdo *funcServerOption) apply(do *serverOptions) {
	fdo.f(do)
}

func newFuncServerOption(f func(*serverOptions)) *funcServerOption {
	return &funcServerOption{
		f: f,
	}
}

func WithName(s string) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.name = s
	})
}
func WithPort(s string) ServerOption {
	return newFuncServerOption(func(o *serverOptions) {
		o.port, _ = strconv.ParseUint(s, 10, 64)
	})
}
