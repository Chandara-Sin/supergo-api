package employee

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getValueForNextSequence(ctx context.Context, db *mongo.Database) (*int, error) {

	collection := getCounterCollection(db)
	filter := bson.M{"collection": "employee"}
	update := bson.M{"$inc": bson.M{"seq_value": 1}}

	if counter, err := collection.CountDocuments(ctx, filter); err != nil {
		return nil, err
	} else if counter == 0 {
		fmt.Println("couter", counter)
		seqDoc := SeqDoc{
			Collection: "employee",
			SeqValue:   0,
		}
		_, err := collection.InsertOne(ctx, seqDoc)
		if err != nil {
			fmt.Println("insert", err.Error())
			return nil, err
		}
	}
	var seqDocRes SeqDoc
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	seqDoc := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	err := seqDoc.Decode(&seqDocRes)

	fmt.Println("err", seqDocRes, err.Error())
	fmt.Println("result", seqDocRes.SeqValue)

	return &seqDocRes.SeqValue, err

	// after := options.After
	// opt := options.FindOneAndUpdateOptions{
	// 	ReturnDocument: &after,
	// }
	// seqDoc, err := collection.UpdateOne(ctx, filter, update)
	// seqDoc := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	// if err != nil {
	// 	return nil, err
	// }
	// if seqDoc.ModifiedCount == 0 {
	// 	return nil, errors.New("cant auto increment")
	// }

	// var seqDocRes SeqDoc
	// rs := collection.FindOne(ctx, filter)
	// err = rs.Decode(&seqDocRes)

	// err := seqDoc.Decode(&rs)
	// fmt.Println(seqDocRes.seqValue)
	// fmt.Println(seqDocRes)

	// return &seqDocRes.seqValue, err
}

func Create(db *mongo.Database) func(context.Context, Employee) (*Employee, error) {
	return func(ctx context.Context, employee Employee) (*Employee, error) {
		collection := getEmployeeCollection(db)
		empId, err := getValueForNextSequence(ctx, db)
		if err != nil {
			return nil, err
		}
		fmt.Println("show", empId)
		reqEmp := Employee{
			EmployeeId: *empId,
			Name:       employee.Name,
			Email:      employee.Email,
			Desc:       employee.Desc,
			CreatedAt:  employee.CreatedAt,
		}
		rs, err := collection.InsertOne(ctx, reqEmp)

		resId := rs.InsertedID.(primitive.ObjectID)
		resEmp := Employee{
			ID:         resId,
			EmployeeId: employee.EmployeeId,
			Name:       employee.Name,
			Email:      employee.Email,
			Desc:       employee.Desc,
			CreatedAt:  employee.CreatedAt,
		}
		return &resEmp, err
	}
}
