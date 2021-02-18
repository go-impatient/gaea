package cache

import (
	"context"
	"time"

	"github.com/Shopify/go-encoding"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var _ Client = (*redisClient)(nil)

type redisClient struct {
	client      *redis.Client
	ExpiresTime time.Duration

	encoding encoding.ValueEncoding
}

func NewRedisClient(c *redis.Client, enc encoding.ValueEncoding) Client {
	return &redisClient{client: c, encoding: enc}
}

func (c *redisClient) GetClient(ctx context.Context) (*redis.Client, error) {
	if c == nil {
		return nil, errors.New("redis client is disable")
	}
	_, err := c.client.Ping(ctx).Result()
	if err != nil {
		return nil, errors.New("new redis client error")
	}
	return c.client, nil
}

func (c *redisClient) Clear(ctx context.Context) error {
	if c == nil {
		return errors.New("redis client is disable")
	}
	client := c.client.WithContext(ctx)
	err := client.FlushDB(ctx).Err()
	return err
}

func (c *redisClient) Get(ctx context.Context, key string, data interface{}) error {
	if c == nil {
		return errors.New("redis cache is disabled")
	}
	client := c.client.WithContext(ctx)
	cmd := client.Get(ctx, key)
	b, err := cmd.Bytes()
	if err != nil {
		if err == redis.Nil {
			return ErrCacheMiss
		}
		return err
	}

	return c.encoding.Decode(b, data)
}

func (c *redisClient) Set(ctx context.Context, key string, data interface{}, expiration time.Time) error {
	if c == nil {
		return errors.New("redis cache is disabled")
	}

	data, err := c.encoding.Encode(data)
	if err != nil {
		return err
	}

	client := c.client.WithContext(ctx)
	cmd := client.Set(ctx, key, data, TtlForExpiration(expiration))
	return cmd.Err()
}

func (c *redisClient) Add(ctx context.Context, key string, data interface{}, expiration time.Time) error {
	if c == nil {
		return errors.New("redis cache is disabled")
	}

	b, err := c.encoding.Encode(data)
	if err != nil {
		return err
	}
	client := c.client.WithContext(ctx)
	cmd := client.SetNX(ctx, key, b, TtlForExpiration(expiration))
	if !cmd.Val() {
		return ErrNotStored
	}
	return cmd.Err()
}

func (c *redisClient) Delete(ctx context.Context, key string) error {
	if c == nil {
		return errors.New("redis cache is disable")
	}

	client := c.client.WithContext(ctx)
	err := client.Del(ctx, key).Err()
	return err
}

func (c *redisClient) Increment(ctx context.Context, key string, delta uint64) (uint64, error) {
	if c == nil {
		return 0, errors.New("redis cache is disable")
	}

	client := c.client.WithContext(ctx)
	cmd := client.IncrBy(ctx, key, int64(delta))
	val, err := cmd.Result()
	return uint64(val), err
}

func (c *redisClient) Decrement(ctx context.Context, key string, delta uint64) (uint64, error) {
	if c == nil {
		return 0, errors.New("redis cache is disable")
	}

	client := c.client.WithContext(ctx)
	cmd := client.DecrBy(ctx, key, int64(delta))
	val, err := cmd.Result()
	return uint64(val), err
}

func (c *redisClient) Exists(ctx context.Context, key string) (bool, error) {
	if c == nil {
		return false, errors.New("redis cache is disable")
	}
	client := c.client.WithContext(ctx)
	count := client.Exists(ctx, key).Val()
	return count > 0, nil
}

func (c *redisClient) Expire(ctx context.Context, key string, expiration time.Duration) error {
	if c == nil {
		return errors.New("redis cache is disable")
	}

	client := c.client.WithContext(ctx)
	err := client.Expire(ctx, key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *redisClient) IncrByExpiration(ctx context.Context, key string, delta uint64, expiration time.Duration) (result uint64, err error) {
	result, err = c.Increment(ctx, key, delta)
	if err != nil {
		return
	}
	err = c.Expire(ctx, key, expiration)
	if err != nil {
		return
	}
	return
}
