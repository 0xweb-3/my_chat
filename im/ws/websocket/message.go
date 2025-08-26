package websocket

type Message struct {
	Method string `json:"method"` // 具体要调用的方法
	FromID string `json:"fromID"` // 消息的请求来源，用于服务方发送到客户方使用
	Data   any    `json:"data"`   // 用户传递的数据
}
