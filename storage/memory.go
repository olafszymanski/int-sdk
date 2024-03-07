package storage

type MemoryStorage struct {
	storage map[string]any
}

func NewMemory() Storager {
	return &MemoryStorage{
		storage: make(map[string]any),
	}
}

func (m *MemoryStorage) StoreAny(key string, value any) error {
	m.storage[key] = value
	return nil
}
