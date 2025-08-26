package websocket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
)

type Server struct {
	patten       string
	sync.RWMutex // 保证连接对象在使用过程中是线程安全的

	routes map[string]HandlerFunc // 记录对应的路由

	// 连接对象的存储
	connToUser     map[*websocket.Conn]string // 实际连接对象到用户
	userToConn     map[string]*websocket.Conn // 用户到连接对象
	authentication Authentication

	addr     string
	upgrader websocket.Upgrader
	logx.Logger
}

func NewServer(addr string, opts ...ServerOptions) *Server {
	opt := newServerOptions(opts...)

	return &Server{
		routes:   make(map[string]HandlerFunc), // 初始化一下
		addr:     addr,
		upgrader: websocket.Upgrader{},

		// 对连接对象初始化
		connToUser: make(map[*websocket.Conn]string),
		userToConn: make(map[string]*websocket.Conn),

		Logger:         logx.WithContext(context.Background()),
		authentication: opt.Authentication,
		patten:         opt.patten,
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	// 处理运行过程中可能抛出的系统性错误
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("server handler ws recover err %v", r)
		}
	}()

	// 得到连接对象
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("upgrade err %v", err)
		return
	}

	// 连接的鉴权
	if !s.authentication.Auth(w, r) {
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("不具备请求权限")))
		return
	}

	// 记录连接通道
	s.addConn(conn, r)

	// 对连接进行处理
	go s.handlerConn(conn)
}

// 根据连接对象执行任务处理
func (s *Server) handlerConn(conn *websocket.Conn) {
	for { // 避免执行一次处理就完毕
		// 获取请求消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			s.Errorf("websocket conn read message err %v", err)
			s.Close(conn)
			return
		}

		//使用goroutine消息是并发处理的 → 可能导致顺序错乱（客户端按顺序发的消息，处理顺序就无法保证）。
		//如果客户端发很多消息，可能会短时间内启动大量 goroutine，占用资源。
		var message Message
		if err = json.Unmarshal(msg, &message); err != nil {
			s.Errorf("json unmarshal err %v, msg %v", err, string(msg))
			s.Close(conn)
			return
		}

		// 根据请求中的method分发路由并执行
		if handler, ok := s.routes[message.Method]; ok {
			handler(s, conn, &message)
		} else {
			// 没有对应处理方式
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("不存在执行的方法%v请检查", message.Method)))
		}
	}
}

// 将路由注册进来
func (s *Server) AddRoutes(rs []Route) {
	for _, r := range rs {
		s.routes[r.Method] = r.Handler
	}
}

func (s *Server) Start() {
	http.HandleFunc(s.patten, s.ServerWs)
	s.Info("启动websocket。。。。。。", s.addr)
	s.Info(http.ListenAndServe(s.addr, nil))
}

func (s *Server) Stop() {
	fmt.Println("停止websocket服务。。。。")
}

// 添加连接对象
func (s *Server) addConn(conn *websocket.Conn, req *http.Request) {
	uid := s.authentication.UserId(req) // 获取用户uid

	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	s.connToUser[conn] = uid
	s.userToConn[uid] = conn
}

func (s *Server) GetConn(uid string) *websocket.Conn {
	s.RWMutex.RLock()
	defer s.RWMutex.Unlock()
	return s.userToConn[uid]
}

func (s *Server) GetConns(uids ...string) []*websocket.Conn {
	if len(uids) == 0 {
		return nil
	}

	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	res := make([]*websocket.Conn, 0, len(uids))
	for _, uid := range uids {
		res = append(res, s.userToConn[uid])
	}
	return res
}

func (s *Server) GetUsers(conns ...*websocket.Conn) []string {
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	var res []string
	if len(conns) == 0 {
		// 获取全部
		res = make([]string, 0, len(s.connToUser))
		for _, uid := range s.connToUser {
			res = append(res, uid)
		}
	} else {
		// 获取部分
		res = make([]string, 0, len(conns))
		for _, conn := range conns {
			res = append(res, s.connToUser[conn])
		}
	}

	return res
}

// 关闭连接
func (s *Server) Close(conn *websocket.Conn) {
	uid := s.connToUser[conn]

	err := conn.Close()
	if err != nil {
		s.Errorf("close websocket conn, uid=%s, err %v", uid, err)
	}
	s.RWMutex.Lock()
	defer s.RWMutex.RUnlock()

	delete(s.connToUser, conn)
	delete(s.userToConn, uid)
}

// 通过用户ID发送消息
func (s *Server) SendByUserId(msg any, sendIds ...string) error {
	if len(sendIds) == 0 {
		return nil
	}

	return s.Send(msg, s.GetConns(sendIds...)...)
}

// 根据连接对象发送消息
func (s *Server) Send(msg any, conns ...*websocket.Conn) error {
	if len(conns) == 0 {
		return nil
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// 投递到每个连接
	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return err
		}
	}
	return nil
}
