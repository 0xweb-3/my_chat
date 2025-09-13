package websocket

import (
	"math"
	"time"
)

const (
	defaultMaxConnectionIdle = time.Duration(math.MaxInt64) // 默认最大连接时长
	defaultAckTimeout        = 30 * time.Second
)
