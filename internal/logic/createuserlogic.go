package logic

import (
	"context"

	"github.com/aaadixu/go-zero-rpc/internal/svc"
	"github.com/aaadixu/go-zero-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.UserCreateRequest) (*user.UserCreateResponse, error) {
	// todo: add your logic here and delete this line

	return &user.UserCreateResponse{
		Id:       312312,
		Username: in.Username,
	}, nil
}
