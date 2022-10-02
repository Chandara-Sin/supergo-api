package counter

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CounterForNextSequence(db *mongo.Database, ctx context.Context, coll string) (int64, error) {
	collection := getCounterCollection(db)

	filter := bson.M{"collection": coll}

	result := Counter{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			counter := Counter{
				Collection: coll,
				SeqValue:   1,
			}

			_, err = collection.InsertOne(ctx, counter)
			return int64(counter.SeqValue), err
		}
		return int64(result.SeqValue), err
	}

	result.SeqValue += 1
	_, err = collection.ReplaceOne(ctx, filter, result)

	return int64(result.SeqValue), err
}
