package storage

type Storager interface {
	GetAny(key string) (any, error)
	StoreAny(key string, value any) error
}
