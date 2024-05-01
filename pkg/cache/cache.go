package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zhikariz/go-commerce/configs"
)

func InitCache(config *configs.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       0,
	})

	return rdb
}

type Cacheable interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
}

type cacheable struct {
	rdb *redis.Client
}

func NewCacheable(rdb *redis.Client) *cacheable {
	return &cacheable{
		rdb: rdb,
	}
}

func (c *cacheable) Set(key string, value interface{}, expiration time.Duration) error {
	operation := c.rdb.Set(context.Background(), key, value, expiration)
	if err := operation.Err(); err != nil {
		return err
	}
	return nil
}

func (c *cacheable) Get(key string) (string, error) {
	val, err := c.rdb.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}
