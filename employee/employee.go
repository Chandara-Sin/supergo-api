package employee

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeId int                `bson:"employee_id" json:"employee_id"`
	Name       string             `bson:"name" json:"name"`
	Email      string             `bson:"email" json:"email"`
	Desc       string             `bson:"desc,omitempty" json:"desc,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

type SeqDoc struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Collection string             `bson:"collection" json:"collection"`
	SeqValue   int                `bson:"seq_value" json:"seq_value"`
}
