package websocket

import "github.com/gorilla/websocket"

type Route struct {
	Method  string      // 执行方法的描述
	Handler HandlerFunc // 具体方法
}

// srv服务对象、连接对象
type HandlerFunc func(srv *Server, conn *websocket.Conn, msg *Message)
