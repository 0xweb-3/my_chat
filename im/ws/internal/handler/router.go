package handler

import (
	"my_chat/im/ws/internal/handler/conversation"
	"my_chat/im/ws/internal/handler/user"
	"my_chat/im/ws/internal/svc"
	"my_chat/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.Online(svc),
		}, {
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
	})
}
