package server

import (
	"flag"
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
	opts      *Options
	rpc       *grpc.Server
	//服务状态信号
	status chan struct{}
}

func NewServer() *Server {
	server := new(Server)
	flag.NewFlagSet("haze", flag.ExitOnError)

	if err := server.loadProperty(); err != nil {
		os.Exit(0)
	}

	server.loadComponent()
	return server
}

//加载属性
func (haze *Server) loadProperty() error {
	haze.opts = &Options{Port: 80}
	haze.opts.discovery=&discovery.DiscoveryOption{DiscoveryType: "nacos"}
	haze.discovery = discovery.NewDiscovery(haze.opts.discovery)
	return nil
}

//初始化组件
func (haze *Server) loadComponent() error {

	return nil
}

func (haze *Server) Start() error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(haze.opts.Port))
	if err != nil {
		log.Fatalf("端口：%v 监听失败 ", haze.opts.Port)
		return err
	}
	s := grpc.NewServer()
	//注册接口
	//test.RegisterWaiterServer(s, &server{})
	/**如果有可以注册多个接口服务,结构体要实现对应的接口方法
	 * user.RegisterLoginServer(s, &server{})
	 * minMovie.RegisterFbiServer(s, &server{})
	 */
	// 在gRPC服务器上注册反射服务
	reflection.Register(s)
	haze.discovery.RegisterInstance(&discovery.Service{})
	haze.rpc = s
	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("服务启动失败，%v ", err)
		return err
	}
	log.Infof("应用启动完成 ")

	return nil
}

func (haze *Server) Shutdown() error {
	return nil
}
