package storage

import (
	"context"
	"sync"
	"time"
)

type item struct {
	value []byte
	ttl   int64
}

type memoryStorage struct {
	storageMutex     sync.RWMutex
	storage          map[string]*item
	hashStorageMutex sync.RWMutex
	hashStorage      map[string]map[string][]byte
}

func NewMemoryStorage() Storager {
	return &memoryStorage{
		storageMutex:     sync.RWMutex{},
		storage:          make(map[string]*item),
		hashStorageMutex: sync.RWMutex{},
		hashStorage:      make(map[string]map[string][]byte),
	}
}

func (s *memoryStorage) Get(_ context.Context, key string) ([]byte, error) {
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
func (s *memoryStorage) Store(_ context.Context, key string, value []byte, expiration time.Duration) error {
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

func (s *memoryStorage) GetHashAll(_ context.Context, hash string) (map[string][]byte, error) {
	s.hashStorageMutex.RLock()
	defer s.hashStorageMutex.RUnlock()

	return s.hashStorage[hash], nil
}

func (s *memoryStorage) StoreHash(_ context.Context, hash string, values map[string]any) error {
	s.hashStorageMutex.Lock()
	defer s.hashStorageMutex.Unlock()

	for k, v := range values {
		s.hashStorage[hash][k] = v.([]byte)
	}
	return nil
}
