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
}
