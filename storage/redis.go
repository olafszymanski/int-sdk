package storage

import (
	"context"
	"fmt"
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

func (s *redisStorage) Get(ctx context.Context, key string) (any, error) {
	r, err := s.client.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%w: key not found", ErrNotFound)
		}
		return nil, err
	}
	return r, nil
}

// If expiration is 0, the key will not expire.
func (s *redisStorage) Store(ctx context.Context, key string, value any, expiration time.Duration) error {
	_, err := s.client.Set(ctx, key, value, expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) GetHashField(ctx context.Context, hash, field string) (any, error) {
	r, err := s.client.HGet(ctx, hash, field).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%w: hash not found", ErrNotFound)
		}
		return nil, err
	}
	return r, nil
}

func (s *redisStorage) GetHashFieldKeys(ctx context.Context, hash string) ([]string, error) {
	res, err := s.client.HKeys(ctx, hash).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%w: hash not found", ErrNotFound)
		}
		return nil, err
	}
	return res, nil
}

func (s *redisStorage) StoreHashField(ctx context.Context, hash string, field string, value any) error {
	_, err := s.client.HSet(ctx, hash, field, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) GetHashFields(ctx context.Context, hash string) (map[string]any, error) {
	res, err := s.client.HGetAll(ctx, hash).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%w: hash not found", ErrNotFound)
		}
		return nil, err
	}
	r := make(map[string]any, len(res))
	for k, v := range res {
		r[k] = any(v)
	}
	return r, nil
}

func (s *redisStorage) StoreHashFields(ctx context.Context, hash string, fields map[string]any) error {
	_, err := s.client.HSet(ctx, hash, fields).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) DeleteHashFields(ctx context.Context, hash string, fields []string) error {
	_, err := s.client.HDel(ctx, hash, fields...).Result()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("%w: hash not found", ErrNotFound)
		}
		return err
	}
	return nil
}

func (s *redisStorage) Close() error {
	return s.client.Close()
}
