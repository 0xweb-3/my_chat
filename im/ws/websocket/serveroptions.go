package websocket

import "time"

// 定义操作

type serverOption struct {
	Authentication
	patten            string
	maxConnectionIdle time.Duration
	ack               AckType       // 设置的ack模式
	ackTimeout        time.Duration // ack应答的等待超时时间
}

type ServerOptions func(opt *serverOption)

func newServerOptions(opts ...ServerOptions) serverOption {
	// 默认值
	o := serverOption{
		Authentication:    new(authentication),
		patten:            "/ws",
		maxConnectionIdle: defaultMaxConnectionIdle,
		ackTimeout:        defaultAckTimeout,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return o
}

func WithServerAuthentication(auth Authentication) ServerOptions {
	return func(opt *serverOption) {
		opt.Authentication = auth
	}
}

func WithServerPatten(patten string) ServerOptions {
	return func(opt *serverOption) {
		opt.patten = patten
	}
}

func WithServerAck(ack AckType) ServerOptions {
	return func(opt *serverOption) {
		opt.ack = ack
	}
}

func WithServerAckTimeout(ackTimeout time.Duration) ServerOptions {
	return func(opt *serverOption) {
		opt.ackTimeout = ackTimeout
	}
}

func WithServerMaxConnectionIdle(maxConnectionIdle time.Duration) ServerOptions {
	return func(opt *serverOption) {
		if maxConnectionIdle > 0 {
			opt.maxConnectionIdle = maxConnectionIdle
		}
	}
}
