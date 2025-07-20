package cache

import (
	"golang.org/x/sync/singleflight"
)

var singleFlight = new(singleflight.Group)

func QueryWithCache[T any](c *Cache, key string, fn func() (T, error)) (T, error) {
	v, err, _ := singleFlight.Do(key, func() (any, error) {
		var res T

		err := c.Get(key, &res)
		if err != nil {
			// 缓存未命中
			res, err = fn()
			if err != nil {
				return nil, err
			}

			if err := c.Set(key, res); err != nil {
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
