package storage

import (
	"fmt"
	"time"
)

var ErrNotFound = fmt.Errorf("key not found")

type item struct {
	value any
	ttl   int64
}

type MemoryStorage struct {
	storage map[string]*item
}

func NewMemoryStorage() Storager {
	return &MemoryStorage{
		storage: make(map[string]*item),
	}
}

func (m *MemoryStorage) GetAny(key string) (any, error) {
	v, ok := m.storage[key]
	if !ok {
		return nil, ErrNotFound
	}
	if time.Now().Unix() < v.ttl {
		delete(m.storage, key)
	}
	return v.value, nil
}

func (m *MemoryStorage) StoreAny(key string, value any, ttl int64) error {
	m.storage[key] = &item{
		value: value,
		ttl:   time.Now().Unix() + ttl,
	}
	return nil
}
