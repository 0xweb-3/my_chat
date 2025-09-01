package user

import (
	"context"
	"github.com/jinzhu/copier"
	"my_chat/user/rpc/user"

	"my_chat/user/api/internal/svc"
	"my_chat/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerResp, err := l.svcCtx.User.Register(l.ctx, &user.RegisterReq{
		Phone:    "",
		Nickname: "",
		Password: "",
		Avatar:   "",
		Sex:      0,
	})
	if err != nil {
		return nil, err
	}

	var res types.RegisterResp
	copier.Copy(&res, registerResp)
	return &res, nil
}
