package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"my_chat/user/api/internal/config"
	"my_chat/user/api/internal/middleware"
)

type ServiceContext struct {
	Config            config.Config
	LoginVerification rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		LoginVerification: middleware.NewLoginVerificationMiddleware().Handle,
	}
}
