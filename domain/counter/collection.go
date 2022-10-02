package counter

import "go.mongodb.org/mongo-driver/mongo"

func getCounterCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("counters")
}
