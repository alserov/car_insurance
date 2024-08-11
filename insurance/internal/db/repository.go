package db

import "context"

type Repository interface {
	Outbox
}

type Outbox interface {
	Create(ctx context.Context, key string) error
	Delete(ctx context.Context, key string) error
}
