package bootstrap

import (
	"fmt"
	"yxkitchen-backend/conf"
	"yxkitchen-backend/pkg/redis"
)

func SetupRedis() {
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", conf.C.Redis.Host, conf.C.Redis.Port),
		conf.C.Redis.Username,
		conf.C.Redis.Password,
		conf.C.Redis.Database,
	)
}
