package logic

import (
	"context"
	"database/sql"
	"fmt"
	"my_chat/demo/userdemo/rpc/models"
	"time"

	"my_chat/demo/userdemo/rpc/internal/svc"
	"my_chat/demo/userdemo/rpc/user"

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

func (l *CreateUserLogic) CreateUser(in *user.CreateUserReq) (*user.CreateUserResp, error) {
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &models.Users{
		Id:       fmt.Sprintf("%v", time.Now().UnixMilli()),
		Avatar:   "Avatar.jpg",
		Name:     fmt.Sprintf("name-%v", time.Now().UnixMilli()),
		Password: sql.NullString{},
	})
	if err != nil {
		l.Errorf("CreateUser error:%v", err)
		return nil, err
	}
	return &user.CreateUserResp{}, nil
}
