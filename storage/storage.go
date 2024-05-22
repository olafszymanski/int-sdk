package storage

import (
	"context"
	"fmt"
	"time"
)

var ErrNotFound = fmt.Errorf("not found")

type Storager interface {
	Get(ctx context.Context, key string) (any, error)
	Store(ctx context.Context, key string, value any, expiration time.Duration) error
	GetHashField(ctx context.Context, hash string, field string) (any, error)
	GetHashFieldKeys(ctx context.Context, hash string) ([]string, error)
	StoreHashField(ctx context.Context, hash string, field string, value any) error
	GetHashFields(ctx context.Context, hash string) (map[string]any, error)
	StoreHashFields(ctx context.Context, hash string, fields map[string]any) error
	DeleteHashFields(ctx context.Context, hash string, fields []string) error
	Close() error
}
