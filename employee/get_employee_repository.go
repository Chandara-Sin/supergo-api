package employee

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEmployeeList(db *mongo.Database) func(context.Context) ([]Employee, error) {
	return func(ctx context.Context) ([]Employee, error) {
		collection := getEmployeeCollection(db)

		//Define an array in which you can store the decoded documents
		var empRes []Employee

		//Passing the bson.D{{}} as the filter matches  documents in the collection
		cur, err := collection.Find(ctx, bson.D{{}})

		for cur.Next(ctx) {
			//Create a value into which the single document can be decoded
			var doc Employee
			err := cur.Decode(&doc)
			if err != nil {
				log.Error(err.Error())
			}
			empRes = append(empRes, doc)
		}

		if err := cur.Err(); err != nil {
			return nil, err
		}

		//Close the cursor once finished
		cur.Close(ctx)

		return empRes, err
	}
}

func GetEmployee(db *mongo.Database) func(context.Context, string) (*Employee, error) {
	return func(ctx context.Context, id string) (*Employee, error) {
		collection := getEmployeeCollection(db)
		var empRes Employee

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		filter := bson.M{"_id": bson.M{"$eq": objId}}
		rs := collection.FindOne(ctx, filter)
		err = rs.Decode(&empRes)

		return &empRes, err
	}
}
