package server

import (
	"github.com/Haze-Lan/haze-go/option"
	"github.com/Haze-Lan/haze-go/provide"
	"google.golang.org/grpc"
	gservice "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
	"sync"
)

var log = grpclog.Component("core")

//每个服务的实列
type Server struct {
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
		log.Errorln(err.Error())
	}
	lis, err := net.Listen("tcp", ":"+strconv.FormatUint(opt.Port, 10))
	if err != nil {
		log.Fatal(err.Error())
	}
	rpc := grpc.NewServer()

	gservice.RegisterChannelzServiceToServer(rpc)
	reflection.Register(rpc)
	server := &Server{
		rpc:  rpc,
		quit: make(chan int),
		lis:  lis,
		opt:  opt,
	}
	return server
}

func (s *Server) Start() error {
	defer s.lis.Close()
	log.Info("应用  %s 启动在本机 %d 完成", s.opt.Name, s.opt.Port)
	//启动grpc
	go func() {
		s.waitGroup.Add(1)
		provide.Register(s.rpc)
		err := s.rpc.Serve(s.lis)
		if err != nil {
			log.Fatal("应用启动失败 ")
		}
		log.Info("grpc 组件  关闭")
		s.waitGroup.Done()
	}()

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
