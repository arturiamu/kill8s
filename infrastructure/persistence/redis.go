package persistence

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"kill8s/config"
	"log"
)

var globalRedis *redis.Client

func InitRedis(config *config.Config) {
	cfg := config.Redis
	globalRedis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password, // 密码
		DB:       cfg.DB,       // 数据库
		PoolSize: cfg.PoolSize, // 连接池大小
	})
	if globalRedis == nil {
		log.Fatalf("err init redis")
	}
}

func GetRedis() *redis.Client {
	return globalRedis
}
