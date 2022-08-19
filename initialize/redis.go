package initialize

import (
	"BookRecSystem/global"
	"context"
	"os"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() *redis.Client {
	redisCfg := global.GSD_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GSD_LOG.Error("Redis Connect Ping Failed, err:", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		global.GSD_LOG.Info("Redis Connect Ping Response:", zap.String("pong", pong))
		return client
	}
}
