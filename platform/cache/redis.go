package cache

import (
	"fmt"
	"strconv"

	"github.com/accalina/restaurant-mgmt/pkg/env"
	"github.com/accalina/restaurant-mgmt/pkg/exception"
	"github.com/go-redis/redis/v8"
)

var RedisCache *redis.Client

func GetRedisCache() *redis.Client {
	return RedisCache
}

func NewRedis() {
	maxPoolSize, err := strconv.Atoi(env.BaseEnv().RedisPoolMaxSize)
	exception.PanicLogging(err)

	minIdlePoolSize, err := strconv.Atoi(env.BaseEnv().RedisPoolMinIdleSize)
	exception.PanicLogging(err)

	RedisCache = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", env.BaseEnv().RedisHost, env.BaseEnv().RedisPort),
		Password:     env.BaseEnv().RedisPass,
		PoolSize:     maxPoolSize,
		MinIdleConns: minIdlePoolSize,
	})
}