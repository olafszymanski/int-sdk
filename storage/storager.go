package storage

import (
	"context"
	"fmt"
	"time"
)

var ErrNotFound = fmt.Errorf("not found")

type Storager interface {
	Get(ctx context.Context, key string) ([]byte, error)
	GetMapValue(ctx context.Context, hash string, key string) ([]byte, error)
	GetMapValues(ctx context.Context, hash string) (map[string][]byte, error)
	GetMapKeys(ctx context.Context, hash string) ([]string, error)

	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	SetMapValue(ctx context.Context, hash string, key string, value any) error
	SetMapValues(ctx context.Context, hash string, values map[string]any) error

	DeleteMapKeys(ctx context.Context, hash string, keys []string) error

	Close() error
}
