package server

import (
	"context"
	"fmt"
	"github.com/Haze-Lan/haze-go/event"
	"github.com/Haze-Lan/haze-go/logger"
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/registry"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
	gservice "google.golang.org/grpc/channelz/service"

	"net"
	"strconv"
	"sync"
)

var log = grpclog.Component("application")

type Server struct {
	//注册中心
	registry  registry.Registry
	rpc       *grpc.Server
	lis       net.Listener
	waitGroup sync.WaitGroup
	opt       *option.ServerOptions
	logger    *grpclog.LoggerV2
	//服务状态信号
	quit chan int
}

func NewServer() *Server {
	logger.NewLogger()
	opt, err := option.LoadServerOptions()
	if err != nil {
		log.Error(err.Error())
	}
	lis, err := net.Listen("tcp", ":"+strconv.FormatUint(opt.Port, 10))
	if err != nil {
		log.Fatal(err.Error())
	}
	rpc := grpc.NewServer()
	gservice.RegisterChannelzServiceToServer(rpc)
	discovery := registry.NewRegistry()
	reflection.Register(rpc)
	server := &Server{
		rpc:      rpc,
		quit:     make(chan int),
		lis:      lis,
		opt:      opt,
		registry: discovery,
	}
	return server
}

func (s *Server) Start() error {
	defer s.Shutdown()
	//启动grpc
	go func() {
		err := s.rpc.Serve(s.lis)
		if err != nil {
			log.Fatal("应用启动失败 ")
		}
		log.Info("grpc 组件  关闭")
	}()
	//注册服务
	go func() {
		var ins = registry.NewInstance(option.WithInstanceAddress(fmt.Sprintf("%s:%d", s.opt.Host, s.opt.Port)),
			option.WithInstanceName(s.opt.Name), option.WithId(fmt.Sprintf("%s:%s:%d", s.opt.Name, s.opt.Host, s.opt.Port)))
		s.registry.RegisterService(context.TODO(), ins)

	}()
	log.Infof("应用[%s]启动在本机[%d]端口完成", s.opt.Name, s.opt.Port)
	select {
	case <-s.quit:
		s.waitGroup.Wait()
		log.Info("应用停止")
	}
	return nil
}

func (s *Server) Shutdown() {
	log.Info("System starts to exit")
	s.rpc.Stop()
	event.GlobalEventBus.Publish(event.EVENT_TOPIC_SERVER_QUIT, nil)
}
