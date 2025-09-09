package msgTransfer

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"my_chat/im/task/mq/internal/handler/svc"
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
func (m *MsgChatTransfer) Consume(ctx context.Context, key string, value string) error {
	fmt.Println("kafka message:", key, "===", value)
	return nil
}
