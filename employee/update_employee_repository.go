package employee

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(db *mongo.Database) func(context.Context, Employee) error {
	return func(ctx context.Context, employee Employee) error {
		collection := getEmployeeCollection(db)

		objID, err := primitive.ObjectIDFromHex(employee.ID.Hex())
		if err != nil {
			return err
		}

		aByte, err := bson.Marshal(employee)
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
			return errors.New("can not update")
		}
		return err
	}
}
