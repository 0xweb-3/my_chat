package ws

import "my_chat/pkg/constants"

// 用于定义定义消息传递过程中的消息格式

// 消息的内容
type Message struct {
	constants.MessageType `mapstructure:"messageType"`
	Content               string `mapstructure:"content"`
}

// 代表聊天
type Chat struct {
	ConversationID string `mapstructure:"conversationID"` // 会话id
	SendId         string `mapstructure:"sendId"`         // 发送方
	ReceiveId      string `mapstructure:"receiveId"`      // 接收方
	Message        `mapstructure:"message"`
	SendTime       int64 `mapstructure:"sendTime"`

	constants.ChatType `mapstructure:"chatType"`
}

// 向kafka中推送的消息结构
type PushMessage struct {
	ConversationID     string `mapstructure:"conversationID"` // 会话id
	SendId             string `mapstructure:"sendId"`         // 发送方
	ReceiveId          string `mapstructure:"receiveId"`      // 接收方
	constants.ChatType `mapstructure:"chatType"`
	SendTime           int64 `mapstructure:"sendTime"`

	constants.MessageType `mapstructure:"messageType"`
	Content               string `mapstructure:"content"`
}
