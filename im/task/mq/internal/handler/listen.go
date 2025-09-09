package handler

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"my_chat/im/task/mq/internal/handler/msgTransfer"
	"my_chat/im/task/mq/internal/handler/svc"
)

// 一个服务支持多个消费者

type Listen struct {
	svc *svc.ServiceContext
}

func NewListen(svc *svc.ServiceContext) *Listen {
	return &Listen{
		svc: svc,
	}
}

// 返回多个消费者
func (l *Listen) Services() []service.Service {
	return []service.Service{
		// 这里可以加载多个消费者
		kq.MustNewQueue(l.svc.Config.MsgChatTransfer, msgTransfer.NewMsgChatTransfer(l.svc)), // 创建出kafka消费者
	}
}
