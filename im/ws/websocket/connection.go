package websocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

type HeartbeatConnection struct {
	idleMu sync.Mutex
	*websocket.Conn
	s                 *Server
	idle              time.Time     // 最后活跃时间
	maxConnectionIdle time.Duration // 最大连接时间
	done              chan struct{}
}

func NewHeartbeatConnection(s *Server, w http.ResponseWriter, r *http.Request) *HeartbeatConnection {
	c, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.Errorf("upgrade err %v", err)
		return nil
	}

	conn := &HeartbeatConnection{
		Conn:              c,
		s:                 s,
		idle:              time.Now(),
		maxConnectionIdle: s.opt.maxConnectionIdle,
		done:              make(chan struct{}),
	}

	// 执行心跳检测
	go conn.keepalive()

	return conn
}

func (c *HeartbeatConnection) keepalive() {
	idleTimer := time.NewTimer(c.maxConnectionIdle)

	defer func() {
		idleTimer.Stop()
	}()

	for {
		select {
		case <-idleTimer.C:
			c.idleMu.Lock()
			idle := c.idle
			if idle.IsZero() { // The connection is non-idle.
				c.idleMu.Unlock()
				idleTimer.Reset(c.maxConnectionIdle)
				continue
			}
			val := c.maxConnectionIdle - time.Since(idle)
			c.idleMu.Unlock()
			if val <= 0 {
				c.s.Close(c)
				return
			}
			idleTimer.Reset(val)
		case <-c.done:
			return
		}
	}
}

func (c *HeartbeatConnection) ReadMessage() (messageType int, p []byte, err error) {
	messageType, p, err = c.Conn.ReadMessage()
	// 非线程安全
	c.idleMu.Lock()
	defer c.idleMu.Unlock()
	c.idle = time.Time{} // 有读操作刷新最后活跃时间
	return
}

func (c *HeartbeatConnection) WriteMessage(messageType int, data []byte) error {
	c.idleMu.Lock()
	// 非线程安全
	defer c.idleMu.Unlock()
	err := c.Conn.WriteMessage(messageType, data)
	c.idle = time.Now() // 有写操作刷新最后活跃时间
	return err
}

func (c *HeartbeatConnection) Close() error {
	select {
	case <-c.done:
	default:
		close(c.done)
	}
	return c.Conn.Close()
}
