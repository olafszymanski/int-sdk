package storage

import (
	"context"
	"fmt"
	"time"
)

var ErrNotFound = fmt.Errorf("not found")

type Storager interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Store(ctx context.Context, key string, value []byte, expiration time.Duration) error
	GetHashField(ctx context.Context, hash string, field string) ([]byte, error)
	GetHashFieldKeys(ctx context.Context, hash string) ([]string, error)
	StoreHashField(ctx context.Context, hash string, field string, value []byte) error
	GetHashFields(ctx context.Context, hash string) (map[string][]byte, error)
	StoreHashFields(ctx context.Context, hash string, fields map[string][]byte) error
	DeleteHashFields(ctx context.Context, hash string, fields []string) error
	Close() error
}
