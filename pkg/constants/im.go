package constants

type MessageType int // 消息类型

const (
	TextMessageType MessageType = iota // 消息类型
)

type ChatType int // 聊天类型
const (
	GroupChatType  ChatType = iota // 群聊类型
	SingleChatType                 // 私聊类型
)
