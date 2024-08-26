package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func MustConnect(addr string) *redis.Client {
	cl := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := cl.Ping(context.Background()).Err(); err != nil {
		panic("failed to ping: " + err.Error())
	}

	return nil
}
