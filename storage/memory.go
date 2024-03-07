package storage

import "fmt"

var ErrNotFound = fmt.Errorf("key not found")

type MemoryStorage struct {
	storage map[string]any
}

func NewMemory() Storager {
	return &MemoryStorage{
		storage: make(map[string]any),
	}
}

func (m *MemoryStorage) GetAny(key string) (any, error) {
	v, ok := m.storage[key]
	if !ok {
		return nil, ErrNotFound
	}
	return v, nil
}

func (m *MemoryStorage) StoreAny(key string, value any) error {
	m.storage[key] = value
	return nil
}
