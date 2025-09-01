package user

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"my_chat/pkg/ctxdata"
	"my_chat/user/rpc/user"

	"my_chat/user/api/internal/svc"
	"my_chat/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	uid := ctxdata.GetUID(l.ctx)
	if uid == "" {
		return nil, fmt.Errorf("uid is empty")
	}

	userResp, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{Id: uid})

	if err != nil {
		return nil, err
	}
	var res types.User
	copier.Copy(&res, userResp.User)
	return &types.UserInfoResp{
		Info: res,
	}, nil
}
