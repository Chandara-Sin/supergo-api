package counter

import "go.mongodb.org/mongo-driver/bson/primitive"

type Counter struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Collection string             `bson:"collection"`
	SeqValue   int                `bson:"seq_value"`
}
