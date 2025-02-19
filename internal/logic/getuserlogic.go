package logic

import (
	"context"

	"github.com/karde88/gozero-ms/user/internal/svc"
	"github.com/karde88/gozero-ms/user/types/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	return &user.UserResponse{
		Id:     in.Id,
		Name:   "phylon",
		Gender: "武装直升机",
	}, nil
}
