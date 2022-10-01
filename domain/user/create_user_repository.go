package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Database) func(context.Context, User) error {
	return func(ctx context.Context, usr User) error {
		collection := getUserCollection(db)
		_, err := collection.InsertOne(ctx, usr)
		return err
	}
}
