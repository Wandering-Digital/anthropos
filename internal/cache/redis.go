package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Wandering-Digital/anthropos/internal/customerror"

	"github.com/redis/go-redis/v9"
	"github.com/thedevsaddam/retry"
)

func NewRedis(client *redis.Client, prefix string, defaultExpiry time.Duration) Cache {
	return &Redis{
		client:        client,
		prefix:        prefix,
		defaultExpiry: defaultExpiry,
	}
}

type Redis struct {
	client        *redis.Client
	prefix        string
	defaultExpiry time.Duration
}

func (r *Redis) BuildKey(keywords ...string) string {

	keywordArray := []string{strings.TrimSuffix(r.prefix, ":")}

	for _, keyword := range keywords {
		keyword = strings.TrimSpace(keyword)
		if keyword != "" && keyword != r.prefix {
			keywordArray = append(keywordArray, keyword)
		}
	}

	key := strings.Join(keywordArray, ":")

	return key
}

func (r *Redis) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *Redis) Set(ctx context.Context, key string, value string, expiration ...time.Duration) error {

	var expiry time.Duration

	if len(expiration) == 0 { // if optional ttl for cache is not provided, read the ttl from consul as default
		expiry = r.defaultExpiry
	} else {
		expiry = expiration[0]
	}

	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	if err := r.client.Set(ctx, key, value, expiry).Err(); err != nil {
		log.Println("error: %v [set redis cache]", err.Error())
		return err
	}

	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {

	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	resStr, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", customerror.ErrCacheNotFound
		}
		return "", err
	}
	return resStr, nil
}

func (r *Redis) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.client.Expire(ctx, key, expiration).Err()
}

func (r *Redis) Keys(ctx context.Context, pattern string) ([]string, error) {
	if !strings.HasPrefix(pattern, r.prefix) {
		pattern = r.BuildKey(pattern)
	}

	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		if err != redis.Nil {
			return nil, err
		}
	}

	if len(keys) < 1 {
		return nil, customerror.ErrCacheNotFound
	}

	return keys, nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	return r.client.Del(ctx, key).Err()
}

func (r *Redis) IncrBy(ctx context.Context, key string, value int64) error {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	return r.client.IncrBy(ctx, key, value).Err()
}

func (r *Redis) DecrBy(ctx context.Context, key string, value int64) error {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	return r.client.DecrBy(ctx, key, value).Err()
}

func (r *Redis) Exists(ctx context.Context, key string) (int64, error) {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	return r.client.Exists(ctx, key).Result()
}

// LRange resolves in positive order of list
func (r *Redis) LRange(ctx context.Context, key string, cursor int, stopCursor ...int) ([]string, error) {
	var response []string
	const incr = 9

	stop := -1
	if stopCursor != nil {
		stop = stopCursor[0]
	}

	if stop < -1 || cursor < -1 {
		return nil, fmt.Errorf("invalid index")
	}

	iter := func(ctx context.Context, cursor int) ([]string, error) {
		if (stop < (cursor+incr) && stop > 0) || cursor == stop {
			result, err := r.client.LRange(ctx, key, int64(cursor), int64(stop)).Result()
			if err != nil {
				return nil, err
			}

			return result, nil
		}

		if cursor < stop || stop == -1 {
			result, err := r.client.LRange(ctx, key, int64(cursor), int64(cursor+incr)).Result()
			if err != nil {
				return nil, err
			}

			return result, nil
		}

		return nil, nil
	}

	for {
		var elements []string

		err := retry.DoFunc(3, 1*time.Second, func() error {
			chunk, err := iter(ctx, cursor)
			if err != nil {
				return err
			}

			elements = chunk

			return nil
		})
		if err != nil {
			return nil, err
		}

		if len(elements) == 0 {
			break
		}

		response = append(response, elements...)

		if cursor != -1 {
			cursor = len(response)
			continue
		}

		break
	}

	return response, nil
}

func (r *Redis) LPush(ctx context.Context, key string, value string) error {
	if !strings.HasPrefix(key, r.prefix) {
		key = r.BuildKey(key)
	}

	return r.client.LPush(ctx, key, value).Err()
}
