package configuration

import (
	"fmt"
	"strconv"

	"github.com/accalina/restaurant-mgmt/exception"
	"github.com/go-redis/redis/v8"
)

var RedisCache *redis.Client

func NewRedis(config Config) *redis.Client {
	host := config.Get("REDIS_HOST")
	port := config.Get("REDIS_PORT")
	pass := config.Get("REDIS_PASS")
	maxPoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MAX_SIZE"))
	exception.PanicLogging(err)

	minIdlePoolSize, err := strconv.Atoi(config.Get("REDIS_POOL_MIN_IDLE_SIZE"))
	exception.PanicLogging(err)

	redisCache := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Password:     pass,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})
	RedisCache = redisCache
	return redisCache
}
