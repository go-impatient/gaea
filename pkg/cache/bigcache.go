package cache

import (
	"context"
	"strconv"
	"sync"
	"time"

	"github.com/allegro/bigcache"
	"github.com/pkg/errors"
)

var _ Client = &bigCacheClient{}

type bigCacheClient struct {
	client *bigcache.BigCache

	encoding Encoding

	mu sync.Mutex
}

func NewBigCacheClient(instance *bigcache.BigCache, encoding Encoding) *bigCacheClient {
	ins := &bigCacheClient{
		client:   instance,
		encoding: encoding,
	}
	return ins
}

func (b *bigCacheClient) Get(ctx context.Context, key string, data interface{}) error {
	if b == nil {
		return errors.New("free xcache is disabled")
	}
	body, err := b.client.Get(key)
	if err != nil && err == bigcache.ErrEntryNotFound {
		return nil
	} else if err != nil {
		return err
	}
	return b.encoding.Decode(body, data)
}

func (b *bigCacheClient) Set(ctx context.Context, key string, data interface{}, expiration time.Time) error {
	if b == nil {
		return errors.New("free xcache is disabled")
	}
	encoded, err := b.encoding.Encode(data)
	if err != nil {
		return err
	}

	return b.client.Set(key, encoded)
}

func (b *bigCacheClient) Add(ctx context.Context, key string, data interface{}, expiration time.Time) error {
	if b == nil {
		return errors.New("free xcache is disabled")
	}
	encoded, err := b.encoding.Encode(data)
	if err != nil {
		return err
	}

	return b.client.Set(key, encoded)
}

func (b *bigCacheClient) Increment(ctx context.Context, key string, delta uint64) (uint64, error) {
	index, err := b.IncrementUint64(ctx, key, delta)
	if err != nil {
		return 0, nil
	}
	return index, nil
}

func (b *bigCacheClient) IncrementUint64(ctx context.Context, key string, delta uint64) (uint64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	exist := true
	body, err := b.client.Get(key)
	if err != nil && err == bigcache.ErrEntryNotFound {
		exist = false
	} else if err != nil {
		return 0, err
	}
	var index uint64 = 0
	if exist {
		n, err := strconv.ParseUint(string(body), 10, 64)
		if err != nil {
			return 0, err
		}
		index = n
	}
	index += delta
	indexSrc := strconv.FormatUint(index, 10)
	err = b.client.Set(key, []byte(indexSrc))
	if err != nil {
		return 0, err
	}
	return index, nil
}

func (b *bigCacheClient) Decrement(ctx context.Context, key string, delta uint64) (uint64, error) {
	index, err := b.DecrementUint64(ctx, key, delta)
	if err != nil {
		return 0, nil
	}
	return index, nil
}

func (b *bigCacheClient) DecrementUint64(ctx context.Context, key string, delta uint64) (uint64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	exist := true
	body, err := b.client.Get(key)
	if err != nil && err == bigcache.ErrEntryNotFound {
		exist = false
	} else if err != nil {
		return 0, err
	}
	var index uint64 = 0
	if exist {
		n, err := strconv.ParseUint(string(body), 10, 64)
		if err != nil {
			return 0, err
		}
		index = n
	}
	index -= delta
	indexSrc := strconv.FormatUint(index, 10)
	err = b.client.Set(key, []byte(indexSrc))
	if err != nil {
		return 0, err
	}
	return index, nil
}

func (b *bigCacheClient) Delete(_ context.Context, key string) error {
	return b.client.Delete(key)
}

func (b *bigCacheClient) Exist(key string) (bool, error) {
	body, err := b.client.Get(key)
	if err != nil {
		return false, err
	}
	return body != nil, nil
}

func (b *bigCacheClient) Iterator() (*bigcache.EntryInfoIterator, error) {
	return b.client.Iterator(), nil
}

func (b *bigCacheClient) Clear(_ context.Context) error {
	return b.client.Reset()
}
