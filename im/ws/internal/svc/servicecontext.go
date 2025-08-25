package svc

import "my_chat/im/ws/internal/config"

// 整个服务使用的上下文

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
