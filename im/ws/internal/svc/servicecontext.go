package svc

import (
	"my_chat/im/immodels"
	"my_chat/im/ws/internal/config"
)

// 整个服务使用的上下文

type ServiceContext struct {
	Config config.Config

	immodels.ChatLogModel // 聊天记录的模型
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ChatLogModel: immodels.NewChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}
