package websocket

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type HeartbeatConnection struct {
	idleMu sync.Mutex
	*websocket.Conn
	s                 *Server
	idle              time.Time     // 最后活跃时间
	maxConnectionIdle time.Duration // 最大连接时间
	done              chan struct{}
	Uid               string

	// 读消息读队列
	messageMu      sync.Mutex
	readMessage    []*Message          // 表示发送过来已经被读取
	readMessageSeq map[string]*Message // 序列化后的被读取的消息
	message        chan *Message       // ack确认消息完成，并将消息交给任务处理的通知

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

		readMessage:    make([]*Message, 0, 2),
		readMessageSeq: make(map[string]*Message, 2),
		message:        make(chan *Message, 1), // 减少阻塞，并保证顺序性
	}

	// 执行心跳检测
	go conn.keepalive()

	return conn
}

// 将消息添加到队列中
func (c *HeartbeatConnection) appendMsgMq(msg *Message) {
	c.messageMu.Lock()
	defer c.messageMu.Unlock()
	// 判断已经存在
	if m, ok := c.readMessageSeq[msg.Id]; ok {
		// 消息记录已经存在，有ack确认过程
		if len(c.readMessage) == 0 {
			// 消息队列里没有消息
			return
		}

		// msg.AckSeq >= m.AckSeq 还没进行ack确认
		if msg.AckSeq >= m.AckSeq {
			// 还没有进行ack确认，重复发送
			return
		}

		// 进行序号的更新
		c.readMessageSeq[msg.Id] = msg
		return
	}

	// 还没有进行ack确认，可能收到多条ack消息
	if msg.FrameType == FrameAck {
		return
	}
	// 将消息记录到队列中
	c.readMessage = append(c.readMessage, msg)
	c.readMessageSeq[msg.Id] = msg
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
