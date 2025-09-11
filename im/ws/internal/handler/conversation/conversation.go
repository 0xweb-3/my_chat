package conversation

import (
	"github.com/mitchellh/mapstructure"
	"my_chat/im/task/mq/mq"
	"my_chat/im/ws/internal/svc"
	"my_chat/im/ws/websocket"
	"my_chat/im/ws/ws"
	"my_chat/pkg/constants"
)

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.HeartbeatConnection, msg *websocket.Message) {
		// 接收聊天信息
		var data *ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn) // 错误消息发送给客户端
			return
		}

		// 消息进行处理
		switch data.ChatType {
		case constants.SingleChatType: // 私聊
			// 推送到kafka中
			err := svc.MqClinet.Push(&mq.MsgChatTransfer{
				ConversationID: data.ConversationID,
				SendId:         data.SendId,
				ReceiveId:      data.ReceiveId,
				ChatType:       data.ChatType,
				SendTime:       data.SendTime,
				MessageType:    data.Message.MessageType,
				Content:        data.Message.Content,
			})
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

		//err := logic.NewConversation(context.Background(), srv, svc).SingleChat(data, conn.Uid)
		//if err != nil {
		//	srv.Send(websocket.NewErrMessage(err), conn)
		//	return
		//}
		//err = srv.SendByUserId(websocket.NewMessage(conn.Uid, ws.Chat{
		//	ConversationID: data.ConversationID,
		//	SendId:         conn.Uid,
		//	ReceiveId:      data.ReceiveId,
		//	Message:        data.Message,
		//	SendTime:       data.SendTime,
		//	ChatType:       data.ChatType,
		//}), data.ReceiveId)
		//
		//if err != nil {
		//	srv.Errorf("Chat err = %v", err)
		//}
		//srv.Send(websocket.NewErrMessage(err), conn)
		//return
		//}

	}
}
