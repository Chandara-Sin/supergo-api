package user

import (
	"Chandara-Sin/supergo-api/domain/counter"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Database) func(context.Context, User) error {
	return func(ctx context.Context, usr User) error {
		collection := getUserCollection(db)

		usrID, err := counter.CounterForNextSequence(db, ctx, "users")
		if err != nil {
			return err
		}

		usr.SetUserID(usrID)
		_, err = collection.InsertOne(ctx, usr)
		return err
	}
}
