package cache

import (
	"testing"

	"github.com/retalkgo/retalk/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestQueryWithCache(t *testing.T) {
	// 创建内存缓存实例
	memCache, err := New(&config.CacheConfig{
		Type: "memory",
		TTL:  1, // 1分钟
	})
	assert.NoError(t, err)

	called := 0
	fn := func() (int, error) {
		called++
		return 12345, nil
	}

	// 第一次调用，缓存未命中，函数执行
	res, err := QueryWithCache(memCache, "test", fn)
	assert.NoError(t, err)
	assert.Equal(t, 12345, res)
	assert.Equal(t, 1, called)

	// 第二次调用，缓存命中，函数不执行
	res2, err := QueryWithCache(memCache, "test", fn)
	assert.NoError(t, err)
	assert.Equal(t, 12345, res2)
	assert.Equal(t, 1, called)
}
