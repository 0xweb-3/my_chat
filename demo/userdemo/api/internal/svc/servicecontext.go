package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"my_chat/demo/userdemo/api/internal/config"
	"my_chat/demo/userdemo/api/internal/middleware"
	"my_chat/demo/userdemo/rpc/userclient"
)

type ServiceContext struct {
	Config            config.Config
	LoginVerification rest.Middleware // 中间件

	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle, // 中间件

		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
