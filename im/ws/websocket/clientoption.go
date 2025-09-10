package websocket

import "net/http"

type ClientOptions func(option *clientOption)

type clientOption struct {
	patten string
	header http.Handler
}

func newClientOptions(opts ...ClientOptions) clientOption {
	o := clientOption{
		patten: "/ws",
		header: nil,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return o
}

func WithClientPatten(patten string) ClientOptions {
	return func(opt *clientOption) {
		opt.patten = patten
	}
}

func WithClientHeader(header http.Header) ClientOptions {
	return func(opt *clientOption) {
		opt.header = header
	}
}
