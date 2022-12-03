package user

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserList(db *mongo.Database) func(context.Context) ([]User, error) {
	return func(ctx context.Context) ([]User, error) {
		collection := getUserCollection(db)

		cur, err := collection.Find(ctx, bson.D{{}})

		rs := []User{}
		for cur.Next(ctx) {
			var doc User
			err := cur.Decode(&doc)
			if err != nil {
				log.Error(err.Error())
			}
			rs = append(rs, doc)
		}

		if err := cur.Err(); err != nil {
			return nil, err
		}

		defer cur.Close(ctx)

		return rs, err
	}
}

func GetUser(db *mongo.Database) func(context.Context, string) (*User, error) {
	return func(ctx context.Context, id string) (*User, error) {
		collection := getUserCollection(db)
		urs := User{}

		usrID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}

		filter := bson.M{"_id": bson.M{"$eq": usrID}}
		err = collection.FindOne(ctx, filter).Decode(&urs)

		return &urs, err
	}
}
