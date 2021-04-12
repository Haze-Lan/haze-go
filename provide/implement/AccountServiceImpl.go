package implement

import (
	"context"
	"github.com/Haze-Lan/haze-go/provide/endpoint"
	"github.com/Haze-Lan/haze-go/provide/model"
)

type AccountServiceImpl struct {
	endpoint.UnimplementedAccountServer
}

func (a AccountServiceImpl) mustEmbedUnimplementedAccountServer() {
	panic("implement me")
}

// Sends a greeting
func (a AccountServiceImpl) Authentication(context.Context, *model.LoginRequest) (*model.LoginResponse, error) {
	return &model.LoginResponse{Message: "哈哈哈"}, nil
}
