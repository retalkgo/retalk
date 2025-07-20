package cache

import (
	"testing"

	"github.com/retalkgo/retalk/internal/config"
	"github.com/stretchr/testify/assert"
)

func Sum(a, b int) int {
	return a + b
}

func Greet(name string) string {
	return "Hello, " + name
}

func DoNothing() {}

type User struct {
	ID   int
	Name string
}

func ProcessUser(u User) string {
	return u.Name
}

func TestGetFuncCallingID(t *testing.T) {
	// Case 1: 调用同一个函数，但使用不同的参数，ID应该不同
	t.Run("Same function, different args", func(t *testing.T) {
		id1 := getFuncCallingID(Sum, 1, 2)
		assert.NotEmpty(t, id1)

		id2 := getFuncCallingID(Sum, 2, 1)
		assert.NotEmpty(t, id2)

		id3 := getFuncCallingID(Sum, 10, 20)
		assert.NotEmpty(t, id3)

		assert.NotEqual(t, id1, id2, "ID should be different for Sum(1,2) and Sum(2,1)")
		assert.NotEqual(t, id1, id3, "ID should be different for Sum(1,2) and Sum(10,20)")
	})

	// Case 2: 多次调用同一个函数且参数完全相同，ID应该相同
	t.Run("Same function, same args", func(t *testing.T) {
		id1 := getFuncCallingID(Greet, "World")
		id2 := getFuncCallingID(Greet, "World")

		assert.NotEmpty(t, id1)
		assert.Equal(t, id1, id2, "ID must be the same for identical function calls")
	})

	// Case 3: 调用不同的函数，ID应该不同
	t.Run("Different functions", func(t *testing.T) {
		id1 := getFuncCallingID(Sum, 1, 2)
		id2 := getFuncCallingID(Greet, "1,2")

		assert.NotEmpty(t, id1)
		assert.NotEmpty(t, id2)
		assert.NotEqual(t, id1, id2, "IDs for different functions (Sum vs Greet) must be different")
	})

	// Case 4: 使用结构体作为参数
	t.Run("Struct as argument", func(t *testing.T) {
		user1 := User{ID: 101, Name: "Alice"}
		user2 := User{ID: 101, Name: "Alice"} // 内容完全相同的另一个实例
		user3 := User{ID: 102, Name: "Bob"}   // 内容不同的实例

		id1 := getFuncCallingID(ProcessUser, user1)
		id2 := getFuncCallingID(ProcessUser, user2)
		id3 := getFuncCallingID(ProcessUser, user3)

		assert.NotEmpty(t, id1)
		assert.Equal(t, id1, id2, "Calls with identical struct values should produce the same ID")
		assert.NotEqual(t, id1, id3, "Calls with different struct values should produce different IDs")
	})

	// Case 5: 调用无参函数
	t.Run("Function with no arguments", func(t *testing.T) {
		id1 := getFuncCallingID(DoNothing)
		id2 := getFuncCallingID(DoNothing)

		assert.NotEmpty(t, id1)
		assert.Equal(t, id1, id2)
		assert.NotContains(t, id1, ":", "ID for a no-arg function should not contain the args separator")
	})

	// Case 6: 错误条件测试 (非常重要!)
	// 现在测试在出错时是否返回空字符串
	t.Run("Error conditions return empty string", func(t *testing.T) {
		// 测试当输入不是一个函数时，是否返回空字符串
		invalidID1 := getFuncCallingID("this is not a function")
		assert.Empty(t, invalidID1, "Should return an empty string when input is not a function")

		// 测试当参数无法被 JSON 序列化时，是否返回空字符串
		invalidID2 := getFuncCallingID(Sum, make(chan int)) // channel类型无法被序列化
		assert.Empty(t, invalidID2, "Should return an empty string when args are not serializable")
	})
}

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
	res, err := QueryWithCache(memCache, fn)
	assert.NoError(t, err)
	assert.Equal(t, 12345, res)
	assert.Equal(t, 1, called)

	// 第二次调用，缓存命中，函数不执行
	res2, err := QueryWithCache(memCache, fn)
	assert.NoError(t, err)
	assert.Equal(t, 12345, res2)
	assert.Equal(t, 1, called)
}
