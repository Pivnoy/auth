package mongodb

import (
	"auth_reg/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Db *mongo.Database
}

func NewMongo(cfg *config.Config) (*Mongo, error) {
	ctx := context.Background()
	cOpts := options.Client().ApplyURI(cfg.MongoURL)
	mClient, err := mongo.Connect(ctx, cOpts)
	if err != nil {
		return nil, err
	}
	err = mClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &Mongo{
		Db: mClient.Database(cfg.MongoDB),
	}, nil
}
