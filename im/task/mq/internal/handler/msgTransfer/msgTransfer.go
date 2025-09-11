package msgTransfer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"my_chat/im/immodels"
	"my_chat/im/task/mq/internal/handler/svc"
	"my_chat/im/task/mq/mq"
	"my_chat/im/ws/websocket"
	"my_chat/pkg/constants"
)

type MsgChatTransfer struct {
	logx.Logger
	svc *svc.ServiceContext
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}

// 定义kafka消费的接口
func (m *MsgChatTransfer) Consume(key string, value string) error {
	fmt.Println("kafka message:", key, "===", value)
	var (
		data mq.MsgChatTransfer
		ctx  = context.Background()
	)
	if err := json.Unmarshal([]byte(value), &data); err != nil {
		return err
	}
	// 记录数据
	if err := m.addChatLog(ctx, &data); err != nil {
		return err
	}

	// 推送消息
	return m.svc.WsClient.Send(websocket.Message{
		FrameType: websocket.FrameData,
		Method:    "push",                    // 消息推送
		FromID:    constants.SYSTEM_ROOT_UID, // 由系统转发
		Data:      data,
	})
}

// 消息的存储
func (m *MsgChatTransfer) addChatLog(ctx context.Context, data *mq.MsgChatTransfer) error {
	chatLog := immodels.ChatLog{
		ConversationID: data.ConversationID,
		SendId:         data.SendId,
		RecvId:         data.ReceiveId,
		MsgFrom:        0,
		ChatType:       data.ChatType,
		MsgType:        data.MessageType,
		MsgContent:     data.Content,
		SendTime:       data.SendTime,
	}
	return m.svc.ChatLogModel.Insert(ctx, &chatLog)
}
