package svc

import (
	"my_chat/im/immodels"
	"my_chat/im/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	immodels.ChatLogModel
	immodels.ConversationModel
	immodels.ConversationsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		ChatLogModel:       immodels.NewChatLogModel(c.Mongo.Url, c.Mongo.Db),
		ConversationModel:  immodels.NewConversationModel(c.Mongo.Url, c.Mongo.Db),
		ConversationsModel: immodels.NewConversationsModel(c.Mongo.Url, c.Mongo.Db),
	}
}
