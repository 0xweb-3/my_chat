package websocket

import "time"

// 定义操作

type serverOption struct {
	Authentication
	patten            string
	maxConnectionIdle time.Duration
}

type ServerOptions func(opt *serverOption)

func newServerOptions(opts ...ServerOptions) serverOption {
	// 默认值
	o := serverOption{
		Authentication:    new(authentication),
		patten:            "/ws",
		maxConnectionIdle: defaultMaxConnectionIdle,
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

func WithServerMaxConnectionIdle(maxConnectionIdle time.Duration) ServerOptions {
	return func(opt *serverOption) {
		if maxConnectionIdle > 0 {
			opt.maxConnectionIdle = maxConnectionIdle
		}
	}
}
