package cache

import (
	"context"
	"encoding/json"
	"github.com/alserov/car_insurance/gateway/internal/utils"
	rd "github.com/go-redis/redis/v8"
	"time"
)

func newRedis(addr string) Cache {
	return &redis{}
}

type redis struct {
	cl *rd.Client
}

func (r redis) Set(ctx context.Context, key string, val any) error {
	b, err := json.Marshal(val)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	r.cl.Set(ctx, key, b, time.Minute)

	return nil
}

func (r redis) Get(ctx context.Context, key string) (any, error) {
	var val any
	if err := json.Unmarshal([]byte(r.cl.Get(ctx, key).Val()), &val); err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	return val, nil
}
