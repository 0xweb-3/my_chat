package websocket

type FrameType uint8

const (
	FrameData  FrameType = 0x0
	FramePing  FrameType = 0x1
	FrameErr   FrameType = 0x9 // 错误类型 用来给前端使用
	FrameAck   FrameType = 0x2 // ack消息类型
	FrameNoAck FrameType = 0x3 // 表示这种场景下消息不用ack确认

	//FrameHeaders      FrameType = 0x1
	//FramePriority     FrameType = 0x2
	//FrameRSTStream    FrameType = 0x3
	//FrameSettings     FrameType = 0x4
	//FramePushPromise  FrameType = 0x5
	//FrameGoAway       FrameType = 0x7
	//FrameWindowUpdate FrameType = 0x8
	//FrameContinuation FrameType = 0x9
)

type Message struct {
	FrameType `json:"frameType"` // 消息类型，是不是心跳消息和普通消息
	Method    string             `json:"method"` // 具体要调用的方法
	FromID    string             `json:"fromID"` // 消息的请求来源，用于服务方发送到客户方使用
	Data      any                `json:"data"`   // 用户传递的数据
}

func NewMessage(fromId string, data any) *Message {
	return &Message{
		FrameType: FrameData, // 默认为普通数据消息
		Method:    "",
		FromID:    fromId,
		Data:      data,
	}
}

func NewErrMessage(err error) *Message {
	return &Message{
		FrameType: FrameErr,
		Data:      err.Error(),
	}
}
