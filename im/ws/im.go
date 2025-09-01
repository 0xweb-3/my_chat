package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"my_chat/im/ws/internal/config"
	"my_chat/im/ws/internal/handler"
	"my_chat/im/ws/internal/svc"
	"my_chat/im/ws/websocket"
)

var configFile = flag.String("f", "im/ws/etc/dev/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 会设置go-zero中的日志，监听相关的处理
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn, websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)))
	defer srv.Stop()

	handler.RegisterHandlers(srv, ctx) // 把所有路由加载进来

	fmt.Println("start websocket server at", c.ListenOn, "....")
	srv.Start()
}
