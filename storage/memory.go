package storage

import (
	"context"
	"time"
)

type item struct {
	value []byte
	ttl   int64
}

type memoryStorage struct {
	storage map[string]*item
}

func NewMemoryStorage() Storager {
	return &memoryStorage{
		storage: make(map[string]*item),
	}
}

func (s *memoryStorage) Get(_ context.Context, key string) ([]byte, error) {
	v, ok := s.storage[key]
	if !ok {
		return nil, ErrNotFound
	}
	if time.Now().Unix() < v.ttl {
		delete(s.storage, key)
	}
	return v.value, nil
}

func (s *memoryStorage) Store(_ context.Context, key string, value []byte, expiration time.Duration) error {
	var ttl int64 = -1
	if expiration != 0 {
		ttl = time.Now().Add(expiration).Unix()
	}
	s.storage[key] = &item{
		value: value,
		ttl:   ttl,
	}
	return nil
}
