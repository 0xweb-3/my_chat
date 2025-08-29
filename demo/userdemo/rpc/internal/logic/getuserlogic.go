package logic

import (
	"context"

	"my_chat/demo/userdemo/rpc/internal/svc"
	"my_chat/demo/userdemo/rpc/user"

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

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserResp, error) {

	return &user.GetUserResp{
		Id:    in.GetId(),
		Name:  "xin",
		Phone: "17612724518",
	}, nil
}
