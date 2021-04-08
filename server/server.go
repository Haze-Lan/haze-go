package server

import (
	"github.com/Haze-Lan/haze-go/discovery"
	"github.com/Haze-Lan/haze-go/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"strconv"
	"sync"
)

var log = logger.LoggerFactory("core")

//每个服务的实列
type Server struct {
	//注册中心
	discovery discovery.Discovery
	rpc       *grpc.Server
	lis       net.Listener
	opt       serverOptions
	waitGroup sync.WaitGroup
	//服务状态信号
	quit chan int
}

//主要用于初始化配置组件
func NewServer(opts ...ServerOption) *Server {
	opts = append(opts, unParseOpts...)
	for _, o := range opts {
		o.apply(defaultServerOptions)
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
		opt:  *defaultServerOptions,
	}
	return server
}

func (s *Server) Start() error {
	//启动grpc
	go func() {
		s.waitGroup.Add(1)
		err := s.rpc.Serve(s.lis)
		if err != nil {
			log.Fatalf("应用 %s 启动失败，%v ", s.opt.name, err)
		}
		log.Infof("grpc 组件  关闭")
		s.waitGroup.Done()
	}()
	log.Infof("应用  %s 启动在本机 %d 完成", s.opt.name, s.opt.port)
	select {
	case <-s.quit:
		s.waitGroup.Wait()
		log.Infof("应用  %s 停止", s.opt.name)
	}
	return nil
}

//停止应用
func (s *Server) Shutdown() error {
	s.rpc.Stop()
	s.quit <- 1
	return nil
}
