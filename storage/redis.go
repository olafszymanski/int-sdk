package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ErrGet       = fmt.Errorf("could not get key")
	ErrStore     = fmt.Errorf("could not store key")
	ErrGetHash   = fmt.Errorf("could not get all hash keys")
	ErrStoreHash = fmt.Errorf("could not store hash key")
)

type redisStorage struct {
	client *redis.Client
}

func NewRedisStorage(address string) Storager {
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
			return nil, fmt.Errorf("%w: %v", ErrNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", ErrGet, err)
	}
	return r, nil
}

// If expiration is 0, the key will not expire.
func (s *redisStorage) Store(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	_, err := s.client.Set(ctx, key, value, expiration).Result()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrStore, err)
	}
	return nil
}

func (s *redisStorage) GetHashAll(ctx context.Context, hash string) (map[string][]byte, error) {
	res, err := s.client.HGetAll(ctx, hash).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("%w: %v", ErrNotFound, err)
		}
		return nil, fmt.Errorf("%w: %v", ErrGetHash, err)
	}
	r := make(map[string][]byte, len(res))
	for k, v := range res {
		r[k] = []byte(v)
	}
	return r, nil
}

// If expiration is 0, the key will not expire.
func (s *redisStorage) StoreHash(ctx context.Context, hash string, values map[string]any) error {
	_, err := s.client.HSet(ctx, hash, values).Result()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrStoreHash, err)
	}
	return nil
}
