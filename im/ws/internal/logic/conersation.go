package logic

import (
	"context"
	"my_chat/im/immodels"
	"my_chat/im/ws/internal/svc"
	"my_chat/im/ws/websocket"
	"my_chat/im/ws/ws"
	"my_chat/pkg/xid"
	"time"
)

type Conversation struct {
	ctx context.Context
	srv *websocket.Server
	svc *svc.ServiceContext
}

func NewConversation(ctx context.Context, srv *websocket.Server, svc *svc.ServiceContext) *Conversation {
	return &Conversation{
		ctx: ctx,
		srv: srv,
		svc: svc,
	}
}

func (c *Conversation) SingleChat(data *ws.Chat, userId string) error {
	if data.ConversationID == "" {
		data.ConversationID = xid.CombineId(userId, data.ReceiveId)
	}

	// 记录消息
	chatLog := immodels.ChatLog{
		ConversationID: data.ConversationID,
		SendId:         userId,
		RecvId:         data.ReceiveId,
		MsgFrom:        0,
		ChatType:       data.ChatType,
		MsgType:        data.MessageType,
		MsgContent:     data.Content,
		SendTime:       time.Now().UnixNano(),
		Status:         0,
		UpdateAt:       time.Now(),
		CreateAt:       time.Now(),
	}
	err := c.svc.ChatLogModel.Insert(c.ctx, &chatLog)
	return err
}
