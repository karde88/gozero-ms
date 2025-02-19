// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.5
// Source: user.proto

package server

import (
	"context"
	"github.com/karde88/gozero-ms/user/internal/logic"
	"github.com/karde88/gozero-ms/user/internal/server"
	"github.com/karde88/gozero-ms/user/internal/svc"
	"github.com/karde88/gozero-ms/user/types/user"

)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) GetUser(ctx context.Context, in *user.IdRequest) (*user.UserResponse, error) {
	l := logic.NewGetUserLogic(ctx, s.svcCtx)
	return l.GetUser(in)
}
