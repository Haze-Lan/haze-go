package server

import (
	"github.com/Haze-Lan/haze-go/logger"
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
	"sync"
)

var log = logger.Factory("core")

//每个服务的实列
type Server struct {
	//注册中心
	discovery registry.Discovery
	rpc       *grpc.Server
	lis       net.Listener
	waitGroup sync.WaitGroup
	opt       *option.ServerOptions
	//服务状态信号
	quit chan int
}

//主要用于初始化配置组件
func NewServer() *Server {
	opt, err := option.LoadServerOptions()
	if err != nil {
		log.Fatal(err.Error())
	}
	lis, err := net.Listen("tcp", ":"+strconv.FormatUint(opt.Port, 10))
	if err != nil {
		log.Fatal(err.Error())
	}
	rpc := grpc.NewServer()
	discovery := registry.NewDiscovery()
	//TODO 注册服务
	reflection.Register(rpc)
	server := &Server{
		discovery: discovery,
		rpc:       rpc,
		quit:      make(chan int),
		lis:       lis,
		opt:       opt,
	}
	return server
}

func (s *Server) Start() error {
	//启动grpc
	go func() {
		s.waitGroup.Add(1)
		err := s.rpc.Serve(s.lis)
		if err != nil {
			log.Fatal("应用启动失败 ")
		}
		log.Info("grpc 组件  关闭")
		s.waitGroup.Done()
	}()
	//注册服务
	go func() {
		s.waitGroup.Add(1)
		var service = registry.NewService(s.opt)
		err := s.discovery.RegisterInstance(service)
		if err != nil {
			log.Error(err.Error())
		}
		s.waitGroup.Done()
	}()
	log.Info("应用  %s 启动在本机 %d 完成", s.opt.Name, s.opt.Port)
	select {
	case <-s.quit:
		s.waitGroup.Wait()
		log.Info("应用停止")
	}
	return nil
}

//停止应用
func (s *Server) Shutdown() error {
	s.rpc.Stop()
	s.quit <- 1
	return nil
}
