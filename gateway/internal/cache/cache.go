package cache

import "context"

type Cache interface {
	Set(ctx context.Context, key string, val any) error
	Get(ctx context.Context, key string) (any, error)
}

const (
	Redis = iota
)

func NewCache(cacheType uint, addr ...string) Cache {
	switch cacheType {
	case Redis:
		return newRedis(addr[0])
	default:
		panic("invalid cache type")
	}
}
