package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStorage struct {
	client *redis.Client
}

func NewRedisStorage(address string) Storager {
	// TODO: fix options
	c := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	return &redisStorage{
		client: c,
	}
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

// If expiration is 0, the key will not expire.
func (s *redisStorage) Store(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	_, err := s.client.Set(ctx, key, value, expiration).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) GetHashField(ctx context.Context, hash, field string) ([]byte, error) {
	r, err := s.client.HGet(ctx, hash, field).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return r, nil
}

func (s *redisStorage) StoreHashField(ctx context.Context, hash string, field string, value []byte) error {
	_, err := s.client.HSet(ctx, hash, field, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func (s *redisStorage) GetHashFields(ctx context.Context, hash string) (map[string][]byte, error) {
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

func (s *redisStorage) StoreHashFields(ctx context.Context, hash string, fields map[string][]byte) error {
	_, err := s.client.HSet(ctx, hash, fields).Result()
	if err != nil {
		return err
	}
	return nil
}
