package redis

import (
	"context"
	"encoding/json"
	"github.com/alserov/car_insurance/contract/internal/db"
	"github.com/alserov/car_insurance/contract/internal/service/models"
	"github.com/alserov/car_insurance/contract/internal/utils"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewOutbox(conn *redis.Client) db.Outbox {
	return &outbox{conn}
}

type outbox struct {
	conn *redis.Client
}

func (o outbox) Create(ctx context.Context, item models.OutboxItem) error {
	b, err := json.Marshal(item)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	o.conn.Set(ctx, item.ID, b, time.Hour*24*3)

	return nil
}

func (o outbox) Get(ctx context.Context, status int, groupID int) ([]models.OutboxItem, error) {
	keys, _, err := o.conn.Scan(ctx, 0, "*", 0).Result()
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	if len(keys) == 0 {
		return nil, nil
	}

	var items []models.OutboxItem
	for _, key := range keys {
		b, err := o.conn.Get(ctx, key).Result()
		if err != nil {
			return nil, utils.NewError(err.Error(), utils.Internal)
		}

		var item models.OutboxItem
		if err = json.Unmarshal([]byte(b), &item); err != nil {
			return nil, utils.NewError(err.Error(), utils.Internal)
		}

		items = append(items, item)
	}

	return items, nil
}

func (o outbox) Delete(ctx context.Context, id string) error {
	o.conn.Del(ctx, id)

	return nil
}
