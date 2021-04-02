package server

import (
	"flag"
	"github.com/Haze-Lan/haze-go/discovery"
	"log"
	"os"
	"sync"
)

//每个服务的实列
type Server struct {

	//日志输出渠道
	loggers []Logger
	//注册中心
	discovery discovery.Client
	//退出信号接收
	quitCh           chan string
	shutdownComplete chan struct{}
	opts             *Options
	optsMu           sync.RWMutex
	mu               sync.Mutex
}
func Init() *Server {
	fs := flag.NewFlagSet("haze", flag.ExitOnError)
	haze := new(Server)

	if err := haze.loadProperty(); err != nil {
		haze.log.(err.Error())
		os.Exit(0)
	}
	haze.loadComponent()
	return haze
}

//加载属性
func (haze *Server) loadProperty() error {
	haze.discovery = &discovery.Etcd{}
	return nil
}

//初始化组件
func (haze *Server) loadComponent() error {
	haze.discovery.Init()
	return nil
}

func (haze *Server) Start() error {

	return nil
}

func (haze *Server) Shutdown() error  {
	return nil
}