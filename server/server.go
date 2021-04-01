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
	 log  *log.Logger
	 //注册中心
	 discovery discovery.Client
	 //退出信号接收
	 quitCh chan string
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
var usageStr = `
Usage: nats-server [options]
Server Options:
    -a, --addr <host>                Bind to host address (default: 0.0.0.0)
    -p, --port <port>                Use port for clients (default: 4222)
    -n, --name <server_name>         Server name (default: auto)
    -P, --pid <file>                 File to store PID
    -m, --http_port <port>           Use port for http monitoring
    -ms,--https_port <port>          Use port for https monitoring
    -c, --config <file>              Configuration file
    -t                               Test configuration and exit
    -sl,--signal <signal>[=<pid>]    Send signal to nats-server process (stop, quit, reopen, reload)
                                     <pid> can be either a PID (e.g. 1) or the path to a PID file (e.g. /var/run/nats-server.pid)
        --client_advertise <string>  Client URL to advertise to other servers
Logging Options:
    -l, --log <file>                 File to redirect log output
    -T, --logtime                    Timestamp log entries (default: true)
    -s, --syslog                     Log to syslog or windows event log
    -r, --remote_syslog <addr>       Syslog server addr (udp://localhost:514)
    -D, --debug                      Enable debugging output
    -V, --trace                      Trace the raw protocol
    -VV                              Verbose trace (traces system account as well)
    -DV                              Debug and trace
    -DVV                             Debug and verbose trace (traces system account as well)

Authorization Options:
        --user <user>                User required for connections
        --pass <password>            Password required for connections
        --auth <token>               Authorization token required for connections
TLS Options:
        --tls                        Enable TLS, do not verify clients (default: false)
        --tlscert <file>             Server certificate file
        --tlskey <file>              Private key for server certificate
        --tlsverify                  Enable TLS, verify client certificates
        --tlscacert <file>           Client certificate CA for verification

Common Options:
    -h, --help                       Show this message
    -v, --version                    Show version
        --help_tls                   TLS help
`
func Init() *Server  {
	fs := flag.NewFlagSet("haze", flag.ExitOnError)
	fs.Usage =usageStr
	haze:=new(Server)


	haze.log= log.New(os.Stderr, "haze", LstdFlags)
	if  err:=haze.loadProperty();err!=nil {
			haze.log.Fatal(err.Error())
	}
	haze.loadComponent()
	return haze
}
//加载属性
func (haze *Server) loadProperty() error {
	haze.discovery=&discovery.Etcd{}
	return nil
}
//初始化组件
func  (haze *Server) loadComponent() error  {
	haze.discovery.Init()
	return nil
}

func (haze *Server) Run() error {

	return nil
}