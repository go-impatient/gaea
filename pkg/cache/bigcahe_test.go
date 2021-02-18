package cache

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/stretchr/testify/assert"
)

type UserInfo struct {
	Id     string    `json:"id" mapstructure:"id"`
	Mobile string    `json:"mobile" mapstructure:"mobile"`
	Name   string    `json:"Name" mapstructure:"Name"`
	Member *UserInfo `json:"member" mapstructure:"member"`
}

func TestBigCacheGetNil(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         24 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	key := "test-key"
	var data string
	err = client.Get(ctx, key, &data)

	fmt.Printf("输出client: %#v\n", err)

	assert.Nil(t, err)
	assert.EqualValues(t, err, nil)
}

func TestCacheSetGet(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         24 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	key := "test-key"
	err = client.Set(ctx, key, "1", time.Now().UTC())
	assert.Nil(t, err)
	var testdata string
	_ = client.Get(ctx, key, &testdata)
	assert.EqualValues(t, testdata, "1")
}

func TestCacheSetGetObj(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         24 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	member := &UserInfo{
		Id:     "202010010000000001",
		Mobile: "13484903810",
		Name:   "小丽",
	}
	userinfo := &UserInfo{
		Id:     "202010010000000002",
		Mobile: "13484903820",
		Name:   "小慧",
		Member: member,
	}

	err = client.Set(ctx, "userinfo", userinfo, time.Now().UTC())
	assert.Nil(t, err)

	var user *UserInfo
	err = client.Get(ctx, "userinfo", &user)
	fmt.Printf("输出client: %#v\n", user)
	assert.Nil(t, err)
	assert.EqualValues(t, user.Id, "202010010000000002")
	assert.EqualValues(t, user.Mobile, "13484903820")
	assert.EqualValues(t, user.Name, "小慧")
	assert.EqualValues(t, user.Member.Id, "202010010000000001")
	assert.EqualValues(t, user.Member.Mobile, "13484903810")
	assert.EqualValues(t, user.Member.Name, "小丽")
}

func TestCacheSetGetArray(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         24 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	var users []UserInfo
	users = append(users, UserInfo{
		Id:     "ID_10000001",
		Mobile: "MOBILE_1000001",
		Name:   "NAME_1000001",
	})
	users = append(users, UserInfo{
		Id:     "ID_10000002",
		Mobile: "MOBILE_1000002",
		Name:   "NAME_1000002",
	})
	users = append(users, UserInfo{
		Id:     "ID_10000003",
		Mobile: "MOBILE_1000003",
		Name:   "NAME_1000003",
	})
	users = append(users, UserInfo{
		Id:     "ID_10000004",
		Mobile: "MOBILE_1000004",
		Name:   "NAME_1000004",
	})
	users = append(users, UserInfo{
		Id:     "ID_10000005",
		Mobile: "MOBILE_1000005",
		Name:   "NAME_1000005",
	})

	err = client.Set(ctx, "users", users, time.Now().UTC())
	assert.Nil(t, err)

	var resultUsers []UserInfo
	err = client.Get(ctx, "users", &resultUsers)
	fmt.Printf("输出client: %#v\n", resultUsers)
	assert.Nil(t, err)
	assert.EqualValues(t, len(resultUsers), len(users))
	assert.EqualValues(t, resultUsers, users)
}

func TestCacheExpired(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         1 * time.Second,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
		CleanWindow:        2 * time.Second,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	err = client.Set(ctx, "test", "this is value", time.Now().UTC())
	assert.Nil(t, err)

	var data string
	testdata := client.Get(ctx, "test", &data)
	fmt.Printf("1. 输出client: %#v\n", data)
	assert.Nil(t, testdata)
	assert.EqualValues(t, data, "this is value")
	time.Sleep(3 * time.Second)

	var data1 string
	testdata1 := client.Get(ctx, "test", &data1)
	fmt.Printf("2. 输出client: %#v\n", data1)
	assert.Nil(t, testdata1)
	assert.EqualValues(t, data1, "")

}

func TestCacheNoExpired(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         1 * time.Second,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
		CleanWindow:        -1 * time.Second,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	err = client.Set(ctx, "test", "this is value", time.Now().UTC())
	assert.Nil(t, err)

	var data string
	for {
		err := client.Get(ctx, "test", &data)
		if err != nil {
			panic(err)
		}
		fmt.Println(data == "this is value")
		time.Sleep(5 * time.Second)
	}
}

func TestCacheIncrement(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         24 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	for i := 1; i <= 1000; i++ {
		index, err := client.Increment(ctx, "test", 10)
		if err != nil {
			panic(err)
		}

		fmt.Printf("输出Increment: %d\n", index)
		assert.EqualValues(t, int(index), i*10)
		time.Sleep(5 * time.Second)
	}
}

func TestIterator(t *testing.T) {
	var ctx = context.Background()
	bigcacheConfig := bigcache.Config{
		Shards:             2048,
		LifeWindow:         24 * time.Hour,
		MaxEntriesInWindow: 1000 * 10 * 64,
		MaxEntrySize:       500,
		Verbose:            true,
		HardMaxCacheSize:   8192,
	}
	bigCacheInstance, err := bigcache.NewBigCache(bigcacheConfig)
	if !assert.NoError(t, err) {
		t.Fatal("fail to initialize big cache")
	}

	client := NewBigCacheClient(bigCacheInstance, DefaultEncoding)

	err = client.Set(ctx, "test1", []byte("test1 value"), time.Now().UTC())
	assert.Nil(t, err)
	var array []UserInfo
	array = append(array, UserInfo{Id: "ID_1001", Mobile: "MOBILE_1001", Name: "NAME_1001"})
	array = append(array, UserInfo{Id: "ID_1002", Mobile: "MOBILE_1002", Name: "NAME_1002"})
	err = client.Set(ctx, "array", array, time.Now().UTC())
	assert.Nil(t, err)

	_, err = client.Increment(ctx, "test3", 1)
	assert.Nil(t, err)
	_, err = client.Increment(ctx, "test3", 1)
	assert.Nil(t, err)
	_, err = client.Increment(ctx, "test3", 1)
	assert.Nil(t, err)
	_, err = client.Increment(ctx, "test3", 1)
	assert.Nil(t, err)

	iterator, err := client.Iterator()
	assert.Nil(t, err)
	for iterator.SetNext() {
		info, err := iterator.Value()
		assert.Nil(t, err)
		fmt.Println(fmt.Sprintf("key : %s ; value : %s", info.Key(), string(info.Value())))
	}
}
