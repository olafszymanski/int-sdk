package storage

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type item struct {
	value any
	ttl   int64
}

type memoryStorage struct {
	storageMutex     sync.RWMutex
	storage          map[string]*item
	hashStorageMutex sync.RWMutex
	hashStorage      map[string]map[string]any
}

func NewMemoryStorage() Storager {
	return &memoryStorage{
		storageMutex:     sync.RWMutex{},
		storage:          make(map[string]*item),
		hashStorageMutex: sync.RWMutex{},
		hashStorage:      make(map[string]map[string]any),
	}
}

func (s *memoryStorage) Get(_ context.Context, key string) (any, error) {
	s.storageMutex.RLock()
	defer s.storageMutex.RUnlock()

	v, ok := s.storage[key]
	if !ok {
		return nil, ErrNotFound
	}
	if time.Now().Unix() < v.ttl {
		delete(s.storage, key)
	}
	return v.value, nil
}

// If expiration is 0, the key will not expire.
func (s *memoryStorage) Store(_ context.Context, key string, value any, expiration time.Duration) error {
	s.storageMutex.Lock()
	defer s.storageMutex.Unlock()

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

func (s *memoryStorage) GetHashField(_ context.Context, hash, field string) (any, error) {
	s.hashStorageMutex.RLock()
	defer s.hashStorageMutex.RUnlock()

	h, ok := s.hashStorage[hash]
	if !ok {
		return nil, fmt.Errorf("%w: hash not found", ErrNotFound)
	}
	if f, ok := h[field]; ok {
		return f, nil
	}
	return nil, fmt.Errorf("%w: field not found", ErrNotFound)
}

func (s *memoryStorage) GetHashFieldKeys(_ context.Context, hash string) ([]string, error) {
	s.hashStorageMutex.RLock()
	defer s.hashStorageMutex.RUnlock()

	h, ok := s.hashStorage[hash]
	if !ok {
		return nil, fmt.Errorf("%w: hash not found", ErrNotFound)
	}
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, fmt.Sprint(k))
	}
	return keys, nil
}

func (s *memoryStorage) StoreHashField(_ context.Context, hash string, field string, value any) error {
	s.hashStorageMutex.Lock()
	defer s.hashStorageMutex.Unlock()

	if _, ok := s.hashStorage[hash]; !ok {
		s.hashStorage[hash] = make(map[string]any, 1)
	}
	s.hashStorage[hash][field] = value
	return nil
}

func (s *memoryStorage) GetHashFields(_ context.Context, hash string) (map[string]any, error) {
	s.hashStorageMutex.RLock()
	defer s.hashStorageMutex.RUnlock()

	h, ok := s.hashStorage[hash]
	if !ok {
		return nil, fmt.Errorf("%w: hash not found", ErrNotFound)
	}
	return h, nil
}

func (s *memoryStorage) StoreHashFields(_ context.Context, hash string, fields map[string]any) error {
	s.hashStorageMutex.Lock()
	defer s.hashStorageMutex.Unlock()

	if _, ok := s.hashStorage[hash]; !ok {
		s.hashStorage[hash] = make(map[string]any, len(fields))
	}
	for k, v := range fields {
		s.hashStorage[hash][k] = v
	}
	return nil
}

func (s *memoryStorage) DeleteHashFields(_ context.Context, hash string, fields []string) error {
	s.hashStorageMutex.Lock()
	defer s.hashStorageMutex.Unlock()

	if _, ok := s.hashStorage[hash]; !ok {
		return fmt.Errorf("%w: hash not found", ErrNotFound)
	}
	for _, field := range fields {
		delete(s.hashStorage[hash], field)
	}
	return nil
}

func (s *memoryStorage) Close() error {
	return nil
}
