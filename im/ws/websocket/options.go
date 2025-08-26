package websocket

// 定义操作

type serverOption struct {
	Authentication
	patten string
}

type ServerOptions func(opt *serverOption)

func newServerOptions(opts ...ServerOptions) serverOption {
	// 默认值
	o := serverOption{
		Authentication: new(authentication),
		patten:         "/ws",
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
