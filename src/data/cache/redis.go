package cache

import (
	"fmt"
	"time"

	"github.com/Amirrezaghanavati/golang-web-api/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:        cfg.Redis.Password,
		DB:              0,
		DialTimeout:     cfg.Redis.DeaiTimeout * time.Second,
		ReadTimeout:     cfg.Redis.DeaiTimeout * time.Second,
		WriteTimeout:    cfg.Redis.DeaiTimeout * time.Second,
		PoolSize:        cfg.Redis.PoolSice,
		PoolTimeout:     cfg.Redis.PoolTimeout,
		ConnMaxIdleTime: 500 * time.Millisecond,
	})
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}
