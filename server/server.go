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
	//日志实列
	log *log.Logger
	//注册中心
	discovery discovery.Client
	//退出信号接收
	quitCh           chan string
	shutdownComplete chan struct{}
	opts             *Options
	optsMu           sync.RWMutex
	mu               sync.Mutex
}

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func Init() *Server {
	fs := flag.NewFlagSet("haze", flag.ExitOnError)
	haze := new(Server)
	haze.log = log.New(os.Stderr, "haze", LstdFlags)
	if err := haze.loadProperty(); err != nil {
		haze.log.Fatal(err.Error())
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

func (haze *Server) Run() error {

	return nil
}
