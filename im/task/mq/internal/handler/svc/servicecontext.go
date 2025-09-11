package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"my_chat/im/immodels"
	"my_chat/im/task/mq/internal/config"
	"my_chat/im/ws/websocket"
	"my_chat/pkg/constants"
	"net/http"
)

type ServiceContext struct {
	Config   config.Config
	WsClient websocket.Client
	*redis.Redis
	immodels.ChatLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config:       c,
		Redis:        redis.MustNewRedis(c.Redisx),
		ChatLogModel: immodels.NewChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}

	token, err := svc.GetSystemToken()
	if err != nil {
		panic(err)
	}
	// 设置websocket的头token
	header := http.Header{}
	header.Set("Authorization", token)
	svc.WsClient = websocket.NewClient(c.Ws.Host, websocket.WithClientHeader(header))
	return svc
}

// 从user rpc服务中获取token
func (svc *ServiceContext) GetSystemToken() (string, error) {
	return svc.Redis.Get(string(constants.ROOT_TOKEN))
}
