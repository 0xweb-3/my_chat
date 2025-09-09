package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"my_chat/im/task/mq/internal/config"
	"my_chat/im/task/mq/internal/handler"
	"my_chat/im/task/mq/internal/handler/svc"
)

var configFile = flag.String("f", "im/task/mq/etc/task.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 会设置go-zero中的日志，监听相关的处理
	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	listen := handler.NewListen(ctx)

	// 服务组对象
	serviceGroup := service.NewServiceGroup()
	for _, s := range listen.Services() {
		serviceGroup.Add(s)
	}
	fmt.Println("starting queue at ....")
	serviceGroup.Start()
}
