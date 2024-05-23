package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	client *redis.Client
}

func NewRedisStorage(ctx context.Context, address, password string) (Storager, error) {
	c := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	if _, err := c.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	return &redisStorage{
		client: c,
	}, nil
}

func (s *redisStorage) Get(ctx context.Context, key string) ([]byte, error) {
	r, err := s.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return r, nil
}

func (s *redisStorage) GetMapValue(ctx context.Context, hash, key string) ([]byte, error) {
	r, err := s.client.HGet(ctx, hash, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return r, nil
}

func (s *redisStorage) GetMapValues(ctx context.Context, hash string) (map[string][]byte, error) {
	res, err := s.client.HGetAll(ctx, hash).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}

	r := make(map[string][]byte, len(res))
	for k, v := range res {
		r[k] = []byte(v)
	}
	return r, nil
}

func (s *redisStorage) GetMapKeys(ctx context.Context, hash string) ([]string, error) {
	res, err := s.client.HKeys(ctx, hash).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return res, nil
}

// If expiration is 0, the key will not expire.
func (s *redisStorage) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	if _, err := s.client.Set(ctx, key, value, expiration).Result(); err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) SetMapValue(ctx context.Context, hash string, key string, value any) error {
	if _, err := s.client.HSet(ctx, hash, key, value).Result(); err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) SetMapValues(ctx context.Context, hash string, values map[string]any) error {
	if _, err := s.client.HSet(ctx, hash, values).Result(); err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) DeleteMapKeys(ctx context.Context, hash string, keys []string) error {
	if _, err := s.client.HDel(ctx, hash, keys...).Result(); err != nil {
		if err == redis.Nil {
			return ErrNotFound
		}
		return err
	}
	return nil
}

func (s *redisStorage) Close() error {
	return s.client.Close()
}
