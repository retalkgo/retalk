package cache

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/stretchr/testify/assert"
)

// Test invalid dest (non-pointer) in Get
func TestGet_NonPointerDest(t *testing.T) {
	conf := &config.CacheConfig{Type: config.CacheTypeMemory, TTL: 1}
	c, err := New(conf)
	assert.NoError(t, err)

	err = c.Set("key", "value")
	assert.NoError(t, err)

	err = c.Get("key", "notPtr")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "dest 必须为指针")
}

func TestMemoryCache_SetGetDeleteClear(t *testing.T) {
	// setup memory cache
	conf := &config.CacheConfig{Type: config.CacheTypeMemory, TTL: 1}
	c, err := New(conf)
	assert.NoError(t, err)

	// test Set
	err = c.Set("foo", 42)
	assert.NoError(t, err)

	// test Get
	var got int
	err = c.Get("foo", &got)
	assert.NoError(t, err)
	assert.Equal(t, 42, got)

	// test Delete
	err = c.Delete("foo")
	assert.NoError(t, err)
	err = c.Get("foo", &got)
	assert.Error(t, err)

	// test Clear
	err = c.Set("bar", "baz")
	assert.NoError(t, err)
	err = c.Clear()
	assert.NoError(t, err)
	err = c.Get("bar", &got)
	assert.Error(t, err)
}

func TestRedisCache_SetGetDeleteClear(t *testing.T) {
	// start miniredis server
	m, err := miniredis.Run()
	assert.NoError(t, err)
	defer m.Close()

	conf := &config.CacheConfig{Type: config.CacheTypeRedis, Addr: m.Addr(), DB: 0, Username: "", Password: "", TTL: 1}
	c, err := New(conf)
	assert.NoError(t, err)

	// test Set
	err = c.Set("alpha", map[string]any{"num": int8(100)})
	assert.NoError(t, err)

	// test Get into struct/map
	var result map[string]any
	err = c.Get("alpha", &result)
	assert.NoError(t, err)
	assert.Equal(t, int8(100), result["num"])

	// test expiration
	m.FastForward(61 * time.Second)
	err = c.Get("alpha", &result)
	assert.Error(t, err)

	// test Delete
	err = c.Set("beta", "xyz")
	assert.NoError(t, err)
	err = c.Delete("beta")
	assert.NoError(t, err)
	err = c.Get("beta", &result)
	assert.Error(t, err)

	// test Clear
	err = c.Set("gamma", 3.14)
	assert.NoError(t, err)
	err = c.Clear()
	assert.NoError(t, err)
	err = c.Get("gamma", &result)
	assert.Error(t, err)
}

func TestRedisClientPingFailure(t *testing.T) {
	conf := &config.CacheConfig{Type: config.CacheTypeRedis, Addr: "127.0.0.1:65000", DB: 0, TTL: 1}
	c, err := New(conf)
	assert.Nil(t, c)
	assert.Error(t, err)
}
