package mq

import "my_chat/pkg/constants"

// 用于存放消息的格式

type MsgChatTransfer struct {
	ConversationID     string `json:"conversationID"` // 会话id
	SendId             string `json:"sendId"`         // 发送方
	ReceiveId          string `json:"receiveId"`      // 接收方
	constants.ChatType `json:"chatType"`
	SendTime           int64 `json:"sendTime"`

	constants.MessageType `json:"messageType"`
	Content               string `json:"content"`
}
