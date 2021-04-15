package impl

import (
	"context"
	"github.com/Haze-Lan/haze-go/examples/simple/haze-common/endpoint"
	"github.com/Haze-Lan/haze-go/examples/simple/haze-common/model"
)

type AccountService struct {
	endpoint.UnimplementedAccountServer
}

func (AccountService) Authentication(ctx context.Context, r *model.LoginRequest) (*model.LoginResponse, error) {

	return &model.LoginResponse{
		Message: r.Name + "------------" + r.Pass,
	}, nil
}
