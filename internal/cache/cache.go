package cache

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/allegro/bigcache/v3"
	lib_cache "github.com/eko/gocache/lib/v4/cache"
	"github.com/eko/gocache/lib/v4/marshaler"
	"github.com/eko/gocache/lib/v4/store"
	bigcache_store "github.com/eko/gocache/store/bigcache/v4"
	redis_store "github.com/eko/gocache/store/redis/v4"
	"github.com/redis/go-redis/v9"
	"github.com/retalkgo/retalk/internal/config"
	"github.com/sirupsen/logrus"
)

type Cache struct {
	ctx       context.Context
	ttl       time.Duration
	instance  *lib_cache.Cache[any]
	marshaler *marshaler.Marshaler
}

func New(conf *config.CacheConfig) (*Cache, error) {
	ctx := context.Background()

	cache := &Cache{
		ctx: ctx,
		ttl: time.Duration(conf.TTL * int(time.Minute)),
	}

	var cacheStore store.StoreInterface

	switch conf.Type {
	case config.CacheTypeMemory:
		bigCacheClient, err := bigcache.New(
			context.Background(),
			bigcache.DefaultConfig(cache.ttl),
		)
		if err != nil {
			return nil, err
		}
		cacheStore = bigcache_store.NewBigcache(bigCacheClient)
	case config.CacheTypeRedis:
		redisClient := redis.NewClient(&redis.Options{
			Addr:     conf.Addr,
			DB:       conf.DB,
			Username: conf.Username,
			Password: conf.Password,
		})

		err := redisClient.Ping(context.Background()).Err()
		if err != nil {
			return nil, err
		}
		cacheStore = redis_store.NewRedis(redisClient, store.WithExpiration(cache.ttl))
	}

	cacheInstance := lib_cache.New[any](cacheStore)

	cache.instance = cacheInstance
	cache.marshaler = marshaler.New(cache.instance)

	return cache, nil
}

func (c *Cache) Get(key string, dest any) error {
	if reflect.ValueOf(dest).Kind() != reflect.Ptr {
		return fmt.Errorf("dest 必须为指针")
	}

	_, err := c.marshaler.Get(c.ctx, key, dest)
	if err != nil {
		logrus.Debugf("[CACHE] Miss: %s", key)
		return err
	}

	logrus.Debugf("[CACHE] Hit: %s", key)

	return nil
}

func (c *Cache) Set(value any, keys ...string) error {
	for _, key := range keys {
		logrus.Debugf("[CACHE] Set: %s", key)
		if err := c.marshaler.Set(c.ctx, key, value); err != nil {
			return err
		}
	}
	return nil
}

func (c *Cache) Delete(keys ...string) error {
	for _, key := range keys {
		logrus.Debugf("[CACHE] Delete: %s", key)
		if err := c.marshaler.Delete(c.ctx, key); err != nil {
			return err
		}
	}
	return nil
}

func (c *Cache) Clear() error {
	return c.marshaler.Clear(c.ctx)
}
