package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteUser(db *mongo.Database) func(context.Context, string) error {
	return func(ctx context.Context, id string) error {
		collection := getUserCollection(db)

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}

		filter := bson.M{"_id": bson.M{"$eq": objId}}
		rs, err := collection.DeleteOne(ctx, filter)

		if rs.DeletedCount == 0 {
			return errors.New("can't delete employee")
		}

		return err
	}
}
