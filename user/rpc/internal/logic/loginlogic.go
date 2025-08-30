package logic

import (
	"context"
	"github.com/pkg/errors"
	"my_chat/pkg/ctxdata"
	"my_chat/pkg/encrypt"
	"my_chat/user/rpc/models"
	"time"

	"my_chat/user/rpc/internal/svc"
	"my_chat/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneIsNotFound = errors.New("用户未注册")
	ErrPasswordError   = errors.New("用户密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 1. 验证用户是否注册，根据手机号
	userEntity, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.GetPhone())
	if err != nil {
		if !errors.Is(err, models.ErrNotFound) {
			return nil, errors.WithStack(ErrPhoneIsNotFound)
		}
		return nil, err
	}
	// 2. 验证密码
	if !encrypt.ValidatePasswordHash(in.GetPassword(), userEntity.Password.String) {
		return nil, errors.WithStack(ErrPasswordError)
	}

	// 生成token
	now := time.Now().Unix()
	token, err := ctxdata.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now,
		l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)

	if err != nil {
		return nil, err
	}
	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
