package core_cache

import (
	"context"
	"fmt"
	"time"

	_ "net/http/pprof"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	RDB *redis.Client
}

type Cache interface {
    Get(ctx context.Context, key string) (string, error)
    Set(ctx context.Context, key string, value string, ttl time.Duration) error
}

func NewRedisClient(cfg Config) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%s", cfg.Host,cfg.Port),
		Password:    cfg.Password,
		Username:        cfg.User,
		DB:         cfg.DB,
		MaxRetries:  cfg.MaxRetries,
		DialTimeout: cfg.DialTimeout,
	})

	pingCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := rdb.Ping(pingCtx).Err(); err != nil {
		return nil, fmt.Errorf("ping redis: %w", err)
	}

	return &RedisClient{ 
		RDB: rdb,
	}, nil
}

func (c *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.RDB.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("cache get: %w", err)
	}

	return val, nil
}

func (c *RedisClient) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	if err := c.RDB.Set(ctx, key, value, ttl).Err(); err != nil {
		return fmt.Errorf("cache set: %w", err)
	}

	return nil
}