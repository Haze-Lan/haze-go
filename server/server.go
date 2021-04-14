package server

import (
	"context"
	"fmt"
	"github.com/Haze-Lan/haze-go/event"
	"github.com/Haze-Lan/haze-go/logger"
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/registry"
	"google.golang.org/grpc"
	gservice "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
)

var log = grpclog.Component("application")

type Server struct {
	//注册中心
	registry registry.Registry
	rpc      *grpc.Server
	opt      *option.ServerOptions
	quit   chan int

}

func NewServer() *Server {
	logger.NewLogger()

	rpc := grpc.NewServer()
	gservice.RegisterChannelzServiceToServer(rpc)
	discovery := registry.NewRegistry()
	reflection.Register(rpc)
	server := &Server{
		rpc:      rpc,
		quit:     make(chan int),
		opt:      option.ServerOptionsInstance,
		registry: discovery,
	}
	return server
}

func (s *Server) Start() error {
	defer s.Shutdown()
	//启动grpc
	go startGrpc(s)
	//注册服务
	go registerThisService(s)
	log.Infof("应用[%s]启动在本机[%d]端口完成", s.opt.Name, s.opt.Port)
	<-s.quit
	return nil
}

func (s *Server) Shutdown() {
	log.Info("System starts to exit")
	s.quit <- 1
	s.rpc.Stop()
	event.GlobalEventBus.Publish(event.EVENT_TOPIC_SERVER_QUIT, nil)
}

//启动rpc服务
func startGrpc(s *Server) {
	defer func() { s.quit<-1 }()
	lis, err := net.Listen("tcp", ":"+strconv.FormatUint(s.opt.Port, 10))
	if err != nil {
		log.Fatal(err.Error())
	}
	err = s.rpc.Serve(lis)
	if err != nil {
		log.Fatal("The GRPC component failed to start ")
	}
	log.Info("The GRPC component is closed")
}

//注册本实例
func registerThisService(s *Server) {
	var ins = registry.NewInstance(fmt.Sprintf("%s:%d", s.opt.Host, s.opt.Port))
	s.registry.RegisterService(context.TODO(), ins)
}
