package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/ibyond/go-start/global"
)

var ctx context.Context

func Redis() {
	redisConfig := global.GstConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.DB,       // use default DB
	})
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		global.GstLog.Error(err)
	} else {
		global.GstLog.Info("redis connect ping response:", pong)
		global.GstRedis = client
	}
}
