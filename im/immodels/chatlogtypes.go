package immodels

import (
	"my_chat/pkg/constants"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var DefaultChatLogLimit int64 = 100

type ChatLog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationID string                `bson:"conversationID"` // 私聊通过两个用户的id可以计算会话id
	SendId         string                `bson:"sendId"`
	RecvId         string                `bson:"recvId"`
	MsgFrom        int                   `bson:"msgFrom"`
	ChatType       constants.ChatType    `bson:"chatType"` // 聊天的类型
	MsgType        constants.MessageType `bson:"msgType"`  // 定义消息的类型
	MsgContent     string                `bson:"msgContent"`
	SendTime       int64                 `bson:"sendTime"`
	Status         int64                 `bson:"status"`

	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
