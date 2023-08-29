package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var SessionCli *redis.Client

func Startup() error {
	opt := &redis.Options{
		Addr:     "redis-service:6379",
		Password: "",
		DB:       0,
	}
	// 连接集群内redis
	SessionCli = redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res := SessionCli.Ping(ctx)

	return res.Err()
}
