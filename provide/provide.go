package provide

import (
	"github.com/Haze-Lan/haze-go/provide/endpoint"
	"github.com/Haze-Lan/haze-go/provide/implement"
	"google.golang.org/grpc"
)

func Register(rpc *grpc.Server) {
	endpoint.RegisterAccountServer(rpc, &implement.AccountServiceImpl{})
}
