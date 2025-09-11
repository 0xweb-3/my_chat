package mqclinet

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"my_chat/im/task/mq/mq"
)

type MessageTransferClient interface {
	Push(msg *mq.MsgChatTransfer) error
}

type messageTransferClient struct {
	pusher *kq.Pusher
}

func NewMessageTransferClient(addr []string, topic string, opts ...kq.PushOption) MessageTransferClient {
	return &messageTransferClient{
		pusher: kq.NewPusher(addr, topic),
	}
}

func (p *messageTransferClient) Push(msg *mq.MsgChatTransfer) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.pusher.Push(context.Background(), string(body))
}
