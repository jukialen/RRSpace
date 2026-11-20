package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"rrspace/backend/internal/config"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg *config.Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr(),
		Password: "",
		DB:       0,
	})
	return &RedisClient{client: client}
}

func (rc *RedisClient) Get(key string) (string, error) {
	return rc.client.Get(rc.ctx(), key).Result()
}

func (rc *RedisClient) Set(key, value string, ttl int) error {
	return rc.client.Set(rc.ctx(), key, value, time.Duration(ttl)*time.Second).Err()
}

func (rc *RedisClient) Del(key string) error {
	return rc.client.Del(rc.ctx(), key).Err()
}

func (rc *RedisClient) Close() error {
	return rc.client.Close()
}

func (rc *RedisClient) ctx() context.Context {
	return context.Background()
}