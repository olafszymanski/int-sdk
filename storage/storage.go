package storage

type Storager interface {
	StoreAny(key string, value any) error
}
