package cache

import (
	"context"
	"math"
	"net"
	"time"

	"github.com/pkg/errors"

	"github.com/bradfitz/gomemcache/memcache"
)

var _ Client = &memcacheClient{}

func NewMemcacheClient(c *memcache.Client, encoding Encoding) *memcacheClient {
	return &memcacheClient{client: c, encoding: encoding}
}

type memcacheClient struct {
	client   *memcache.Client
	encoding Encoding
}

// Clear clear all cached in memcache.
func (c *memcacheClient) Clear(_ context.Context) error {
	return c.client.FlushAll()
}

func (c *memcacheClient) Get(_ context.Context, key string, data interface{}) error {
	mItem, err := c.client.Get(key)
	if err != nil {
		// Abstract the memcache-specific error
		if err == memcache.ErrCacheMiss {
			return ErrCacheMiss
		}
		return coalesceTimeoutError(err)
	}

	return c.encoding.Decode(mItem.Value, data)
}

// GetMulti get value from memcache.
func (c *memcacheClient) GetMulti(keys []string) []interface{} {
	size := len(keys)
	var rv []interface{}
	mv, err := c.client.GetMulti(keys)
	if err == nil {
		for _, v := range mv {
			rv = append(rv, v.Value)
		}
		return rv
	}
	for i := 0; i < size; i++ {
		rv = append(rv, err)
	}
	return rv
}

func (c *memcacheClient) Set(_ context.Context, key string, data interface{}, expiration time.Time) error {
	mItem, err := c.encodeItem(key, data, expiration)
	if err != nil {
		return err
	}
	return coalesceTimeoutError(c.client.Set(mItem))
}

func (c *memcacheClient) Add(_ context.Context, key string, data interface{}, expiration time.Time) error {
	mItem, err := c.encodeItem(key, data, expiration)
	if err != nil {
		return err
	}
	err = c.client.Add(mItem)

	if err == memcache.ErrNotStored {
		// Abstract the memcache-specific error
		return ErrNotStored
	}
	return coalesceTimeoutError(err)
}

func (c *memcacheClient) Replace(key string, value interface{}, expires time.Duration) error {
	return c.Replace(key, value, expires)
}

func (c *memcacheClient) Delete(_ context.Context, key string) error {
	err := c.client.Delete(key)
	if err == memcache.ErrCacheMiss {
		// Deleting a missing entry is not an actual issue.
		return nil
	}
	return coalesceTimeoutError(err)
}

func (c *memcacheClient) Increment(ctx context.Context, key string, delta uint64) (uint64, error) {
	newValue, err := c.client.Increment(key, delta)
	if err == memcache.ErrCacheMiss {
		// Initialize
		err = c.Add(ctx, key, delta, NeverExpire)
		if err == ErrNotStored {
			// Race condition, try increment again
			return c.Increment(ctx, key, delta)
		}
		newValue = delta
	}
	return newValue, coalesceTimeoutError(err)
}

func (c *memcacheClient) Decrement(ctx context.Context, key string, delta uint64) (uint64, error) {
	newValue, err := c.client.Decrement(key, delta)
	if err == memcache.ErrCacheMiss {
		// Initialize
		err = c.Add(ctx, key, -delta, NeverExpire)
		if err == ErrNotStored {
			// Race condition, try increment again
			return c.Decrement(ctx, key, delta)
		}
		newValue = -delta
	}
	return newValue, coalesceTimeoutError(err)
}

func (c *memcacheClient) Flush() error {
	return errors.New("cache: not support.")
}

func (c *memcacheClient) encodeItem(key string, data interface{}, expiration time.Time) (*memcache.Item, error) {
	encoded, err := c.encoding.Encode(data)
	if err != nil {
		return nil, err
	}

	mItem := &memcache.Item{
		Value: encoded,
		Key:   key,
	}
	if ttl := TtlForExpiration(expiration); ttl != 0 {
		mItem.Expiration = int32(math.Max(ttl.Seconds(), 1))
	}

	return mItem, nil
}

type connectTimeoutError struct{}

func (connectTimeoutError) Error() string   { return "memcache: connect timeout" }
func (connectTimeoutError) Timeout() bool   { return true }
func (connectTimeoutError) Temporary() bool { return true }

func coalesceTimeoutError(err error) error {
	// For some reason, gomemcache decided to replace the standard net.Error.
	// Coalesce into a generic net.Error so that client don't have to deal with memcache-specific errorsx.
	switch typed := err.(type) {
	case *memcache.ConnectTimeoutError:
		return &net.OpError{
			Err:  &connectTimeoutError{},
			Addr: typed.Addr,
			Net:  typed.Addr.Network(),
			Op:   "connect",
		}
	default:
		// This also work if err is nil
		return err
	}
}
