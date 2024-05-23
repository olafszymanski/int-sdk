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
	storageLock    sync.RWMutex
	storage        map[string]*item
	mapStorageLock sync.RWMutex
	mapStorage     map[string]map[string]any
}

func NewMemoryStorage() Storager {
	return &memoryStorage{
		storageLock:    sync.RWMutex{},
		storage:        make(map[string]*item),
		mapStorageLock: sync.RWMutex{},
		mapStorage:     make(map[string]map[string]any),
	}
}

func (s *memoryStorage) Get(_ context.Context, key string) ([]byte, error) {
	s.storageLock.RLock()
	defer s.storageLock.RUnlock()

	v, ok := s.storage[key]
	if !ok {
		return nil, ErrNotFound
	}
	if time.Now().Unix() < v.ttl {
		delete(s.storage, key)
	}
	return v.value.([]byte), nil
}

func (s *memoryStorage) GetMapValue(_ context.Context, hash, key string) ([]byte, error) {
	s.mapStorageLock.RLock()
	defer s.mapStorageLock.RUnlock()

	h, ok := s.mapStorage[hash]
	if !ok {
		return nil, ErrNotFound
	}
	if v, ok := h[key]; ok {
		return v.([]byte), nil
	}
	return nil, fmt.Errorf("%w: field not found", ErrNotFound)
}

func (s *memoryStorage) GetMapValues(_ context.Context, hash string) (map[string][]byte, error) {
	s.mapStorageLock.RLock()
	defer s.mapStorageLock.RUnlock()

	h, ok := s.mapStorage[hash]
	if !ok {
		return nil, ErrNotFound
	}
	res := make(map[string][]byte, len(h))
	for k, v := range h {
		res[k] = v.([]byte)
	}
	return res, nil
}

func (s *memoryStorage) GetMapKeys(_ context.Context, hash string) ([]string, error) {
	s.mapStorageLock.RLock()
	defer s.mapStorageLock.RUnlock()

	h, ok := s.mapStorage[hash]
	if !ok {
		return nil, ErrNotFound
	}
	keys := make([]string, 0, len(h))
	for k := range h {
		keys = append(keys, fmt.Sprint(k))
	}
	return keys, nil
}

// If expiration is 0, the key will not expire.
func (s *memoryStorage) Set(_ context.Context, key string, value any, expiration time.Duration) error {
	s.storageLock.Lock()
	defer s.storageLock.Unlock()

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

func (s *memoryStorage) SetMapValue(_ context.Context, hash string, key string, value any) error {
	s.mapStorageLock.Lock()
	defer s.mapStorageLock.Unlock()

	if _, ok := s.mapStorage[hash]; !ok {
		s.mapStorage[hash] = make(map[string]any, 1)
	}
	s.mapStorage[hash][key] = value
	return nil
}

func (s *memoryStorage) SetMapValues(_ context.Context, hash string, values map[string]any) error {
	s.mapStorageLock.Lock()
	defer s.mapStorageLock.Unlock()

	if _, ok := s.mapStorage[hash]; !ok {
		s.mapStorage[hash] = make(map[string]any, len(values))
	}
	for k, v := range values {
		s.mapStorage[hash][k] = v
	}
	return nil
}

func (s *memoryStorage) DeleteMapKeys(_ context.Context, hash string, keys []string) error {
	s.mapStorageLock.Lock()
	defer s.mapStorageLock.Unlock()

	if _, ok := s.mapStorage[hash]; !ok {
		return ErrNotFound
	}
	for _, field := range keys {
		delete(s.mapStorage[hash], field)
	}
	return nil
}

func (s *memoryStorage) Close() error {
	return nil
}
