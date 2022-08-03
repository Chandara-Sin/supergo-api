package employee

import "go.mongodb.org/mongo-driver/mongo"

func getEmployeeCollection(db *mongo.Database) *mongo.Collection {
	return db.Collection("employees")
}
