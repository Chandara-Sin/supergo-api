package employee

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Database) func(context.Context, Employee) (primitive.ObjectID, error) {
	return func(ctx context.Context, employee Employee) (primitive.ObjectID, error) {
		collection := getEmployeeCollection(db)
		rs, err := collection.InsertOne(ctx, employee)
		str := rs.InsertedID.(primitive.ObjectID)
		return str, err
	}
}
