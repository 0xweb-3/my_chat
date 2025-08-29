package user

import (
	"github.com/gorilla/websocket"
	"my_chat/im/ws/internal/svc"
	websocketx "my_chat/im/ws/websocket"
)

// 获取所有在线用户
func Online(svc *svc.ServiceContext) websocketx.HandlerFunc {
	return func(srv *websocketx.Server, conn *websocket.Conn, msg *websocketx.Message) {
		uids := srv.GetUsers()
		u := srv.GetUsers(conn) // 该连接对应的用户
		err := srv.Send(websocketx.NewMessage(u[0], uids), conn)
		if err != nil {
			srv.Info("err", err)
		}
	}
}
