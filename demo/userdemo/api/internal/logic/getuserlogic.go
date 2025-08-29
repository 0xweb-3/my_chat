package logic

import (
	"context"
	"my_chat/demo/userdemo/rpc/userclient"

	"my_chat/demo/userdemo/api/internal/svc"
	"my_chat/demo/userdemo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.UserReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.GetUser(l.ctx, &userclient.GetUserReq{
		Id: "122",
	})
	if err != nil {
		l.Error(err)
		return nil, err
	}
	return &types.UserResp{
		Id:    user.GetId(),
		Name:  user.GetName(),
		Phone: user.GetPhone(),
	}, nil
}
