package logic

import (
	"context"

	"github.com/aaadixu/go-zero-rpc/internal/svc"
	"github.com/aaadixu/go-zero-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.UserQueryRequest) (*user.UserQueryResponse, error) {

	return &user.UserQueryResponse{
		Id:       in.Id,
		Username: "user_" + string(in.Id),
		Age:      18,
		Gender: func() string {
			if in.Id%2 == 0 {
				return "男"
			} else {
				return "女"
			}
		}(),
	}, nil
}
