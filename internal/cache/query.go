package cache

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"

	"golang.org/x/sync/singleflight"
)

var singleFlight = new(singleflight.Group)

func getFuncSignatureID(fn any) (string, error) {
	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		return "", fmt.Errorf("input is not a function")
	}

	// 获取函数的完整名称（包括包路径）
	funcInfo := runtime.FuncForPC(val.Pointer())
	if funcInfo == nil {
		return "", fmt.Errorf("cannot get function info")
	}
	fullName := funcInfo.Name() // e.g., "main.Sum"

	// 获取所有输入参数的类型名称
	t := val.Type()
	paramTypes := make([]string, 0, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		paramTypes = append(paramTypes, t.In(i).String())
	}

	// 拼接成最终的签名ID
	return fmt.Sprintf("%s(%s)", fullName, strings.Join(paramTypes, ",")), nil
}

// 根据函数签名和传入的具体参数值生成一个唯一的调用ID
func getFuncCallingID(fn any, args ...any) string {
	// 获取函数签名作为ID的基础部分
	signatureID, err := getFuncSignatureID(fn)
	if err != nil {
		return ""
	}

	// 如果没有参数，调用ID就等于签名ID
	if len(args) == 0 {
		return signatureID
	}

	// 将实际参数值序列化为JSON字符串
	argsJSON, err := json.Marshal(args)
	if err != nil {
		// 如果参数无法被序列化（例如包含channel或func），则返回错误
		return ""
	}

	// 组合签名ID和参数ID
	// 格式: package.Func(type1,type2):[arg1_json,arg2_json,...]
	return fmt.Sprintf("%s:%s", signatureID, string(argsJSON))
}

func QueryWithCache[T any](c *Cache, fn func() (T, error)) (T, error) {
	v, err, _ := singleFlight.Do(getFuncCallingID(fn), func() (any, error) {
		var res T

		err := c.Get(getFuncCallingID(fn), &res)
		if err != nil {
			// 缓存未命中
			res, err = fn()
			if err != nil {
				return nil, err
			}

			if err := c.Set(getFuncCallingID(fn), res); err != nil {
				return nil, err
			}
		}

		return res, nil
	})
	if err != nil {
		return *new(T), err
	}

	return v.(T), nil
}
