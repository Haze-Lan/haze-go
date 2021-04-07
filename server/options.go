package server

import (
	"flag"
	"github.com/Haze-Lan/haze-go/utils"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"reflect"
	"strconv"
)

var defaultServerOptions = serverOptions{}

func init() {
	var configPath string
	base, _ := utils.HomeDir()
	flag.StringVar(&configPath, "config", base+"\\application.yaml", "config file path")
	flag.Parse()
	log.Infoln("加载配置文件：", configPath)
	m := make(map[string]interface{})
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("配置文件读取错误: ", err.Error())
	}
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	for k, v := range m["application"].((map[interface{}]interface{})) {
		fv := reflect.ValueOf("With" + k.(string))
		if fv.Kind() != reflect.Func {
			log.Errorln("参数不能解析：", "application."+k.(string))
		}
		fv.Call([]reflect.Value{reflect.ValueOf(v)})
	}
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
