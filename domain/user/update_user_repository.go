package user

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(db *mongo.Database) func(context.Context, User) error {
	return func(ctx context.Context, usr User) error {
		collection := getUserCollection(db)

		objID, err := primitive.ObjectIDFromHex(usr.ID.Hex())
		if err != nil {
			return err
		}

		aByte, err := bson.Marshal(usr)
		if err != nil {
			return err
		}

		var updatedEmp bson.M
		err = bson.Unmarshal(aByte, &updatedEmp)
		if err != nil {
			return err
		}

		rs, err := collection.UpdateByID(ctx, objID, bson.D{{Key: "$set", Value: updatedEmp}})
		if rs.ModifiedCount == 0 {
			return errors.New("endpoint is not found")
		}
		return err
	}
}
