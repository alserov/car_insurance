package mongo

import (
	"context"
	"github.com/alserov/car_insurance/insurance/internal/db"
	"github.com/alserov/car_insurance/insurance/internal/service/models"
	"github.com/alserov/car_insurance/insurance/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName     = "outbox"
	collection = "outbox"
)

func NewOutbox(conn *mongo.Client) db.Outbox {
	coll := conn.Database(dbName).Collection(collection)

	return &outbox{
		db: coll,
	}
}

type outbox struct {
	db *mongo.Collection
}

func (o outbox) Create(ctx context.Context, item models.OutboxItem) error {
	_, err := o.db.InsertOne(ctx, item)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}

func (o outbox) Get(ctx context.Context, status int, groupID int) ([]models.OutboxItem, error) {
	filter := bson.M{"groupID": groupID, "status": status}

	curs, err := o.db.Find(ctx, filter)
	if err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	var items []models.OutboxItem
	if err = curs.All(ctx, &items); err != nil {
		return nil, utils.NewError(err.Error(), utils.Internal)
	}

	return items, nil
}

func (o outbox) Delete(ctx context.Context, status int, groupID int) error {
	filter := bson.M{"groupID": groupID, "status": status}

	_, err := o.db.DeleteMany(ctx, filter)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}
