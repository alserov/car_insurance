package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MustConnect(addr string) *mongo.Client {
	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(addr))
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	if err = conn.Ping(context.Background(), nil); err != nil {
		panic("failed to ping: " + err.Error())
	}

	return conn
}
