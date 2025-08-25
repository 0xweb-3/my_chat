package config

import "github.com/zeromicro/go-zero/core/service"

type Config struct {
	service.ServiceConf // go-zero中的配置

	ListenOn string
}
