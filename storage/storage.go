package storage

import (
	"context"
	"fmt"
	"time"
)

var ErrNotFound = fmt.Errorf("key not found")

type Storager interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Store(ctx context.Context, key string, value []byte, expiration time.Duration) error
	GetHashAll(ctx context.Context, hash string) (map[string][]byte, error)
	StoreHash(ctx context.Context, hash string, values map[string]any) error
}
