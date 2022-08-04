package employee

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEmployee(db *mongo.Database) func(context.Context, string) (*Employee, error) {
	return func(ctx context.Context, Id string) (*Employee, error) {
		collect := getEmployeeCollection(db)
		var empRes Employee

		objId, err := primitive.ObjectIDFromHex(Id)
		if err != nil {
			return nil, err
		}

		filter := bson.M{"_id": bson.M{"$eq": objId}}
		rs := collect.FindOne(ctx, filter)
		err = rs.Decode(&empRes)

		return &empRes, err
	}
}
