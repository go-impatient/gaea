package cache

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/Shopify/go-encoding"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rc Client
var once sync.Once

var ctx = context.Background()

func setup() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       1,
		Password: "",
	})
	rc = NewRedisClient(client, DefaultEncoding)
}

func TestSetGet(t *testing.T) {
	once.Do(setup)
	err := rc.Set(ctx, "test-src", "1", time.Now().UTC())
	if err != nil {
		t.Errorf("储存数据失败: %s", err)
		return
	}
	assert.Nil(t, err)

	var data string
	res := rc.Get(ctx, "test-src", &data)
	if res != nil {
		t.Errorf("获取数据失败: %s", res)
		return
	}
	assert.Nil(t, res)
	require.Equal(t, data, data)
}

func ExampleNewRedisClient() {
	opts, err := redis.ParseURL("redis://:qwerty@localhost:6379/1")
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opts)
	NewRedisClient(client, DefaultEncoding)
}

func testRedis(t *testing.T) *redis.Client {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}

	url := os.Getenv("REDIS_URL")
	if len(url) == 0 {
		t.Skip("redis client not configured")
		return nil
	}

	opts, err := redis.ParseURL(url)
	require.NoError(t, err)
	return redis.NewClient(opts)
}

func Test_redisClient(t *testing.T) {
	client := testRedis(t)
	encodings := map[string]encoding.ValueEncoding{
		"gob":          gobEncoding,
		"json":         encoding.JSONEncoding,
		"literal+gob":  encoding.NewLiteralEncoding(gobEncoding),
		"literal+json": encoding.NewLiteralEncoding(encoding.JSONEncoding),
	}
	for name, enc := range encodings {
		t.Run(name, func(t *testing.T) {
			testClient(t, NewRedisClient(client, enc), enc)
		})
	}

}
