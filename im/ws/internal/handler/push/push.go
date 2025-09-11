package push

import (
	"github.com/mitchellh/mapstructure"
	"my_chat/im/ws/internal/svc"
	"my_chat/im/ws/websocket"
	"my_chat/im/ws/ws"
)

// 用于向kafka中转发并，推送消息
func Push(ctx *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.HeartbeatConnection, msg *websocket.Message) {
		var data ws.PushMessage
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err))
			return
		}

		// 发送的目标
		receiveConn := srv.GetConn(data.ReceiveId)
		// 说明是离线状态
		if receiveConn == nil {
			// todo 处理目标离线情况
		}

		srv.Infof("Push msg to online user message=%v", data)
		srv.Send(websocket.NewMessage(data.SendId, &ws.Chat{
			ConversationID: data.ConversationID,
			Message: ws.Message{
				MessageType: data.MessageType,
				Content:     data.Content,
			},
			SendTime: data.SendTime,
			ChatType: data.ChatType,
		}), receiveConn)
	}
}
