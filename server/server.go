package server

import (
	"github.com/Haze-Lan/haze-go/discovery"
	"github.com/Haze-Lan/haze-go/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"strconv"
)

var log = logger.LoggerFactory("core")

//每个服务的实列
type Server struct {
	//注册中心
	discovery discovery.Discovery
	rpc       *grpc.Server
	lis       net.Listener
	opt       serverOptions
	//服务状态信号
	quit chan int
}

//主要用于初始化配置组件
func NewServer(opts ...ServerOption) *Server {
	for _, o := range opts {
		o.apply(&defaultServerOptions)
	}
	lis, err := net.Listen("tcp", ":"+strconv.FormatUint(defaultServerOptions.port, 10))
	if err != nil {
		log.Fatalln(err.Error())
	}
	rpc := grpc.NewServer()
	//TODO 注册服务
	reflection.Register(rpc)
	server := &Server{
		rpc:  rpc,
		quit: make(chan int),
		lis:  lis,
		opt:  defaultServerOptions,
	}
	return server
}

func (s *Server) Start() error {
	go func() {
		err := s.rpc.Serve(s.lis)
		if err != nil {
			log.Fatalf("应用 %s 启动失败，%v ", s.opt.name, err)
			os.Exit(1)
		}
	}()
	log.Infof("应用  %s 启动完成", s.opt.name)
	select {
	case <-s.quit:
		log.Infof("应用  %s 停止", s.opt.name)
	}
	return nil
}

//停止应用
func (s *Server) Shutdown() error {
	s.quit <- 1
	return nil
}
