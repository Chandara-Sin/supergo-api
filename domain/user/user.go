package user

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Gender string

const (
	Male          Gender = "male"
	Female        Gender = "female"
	NotApplicable Gender = "n/a"
)

func (usr *User) SetUserID(usrId int64) {
	usr.UserId = "US" + fmt.Sprintf("%05d", usrId)
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId    string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
	FirstName string             `bson:"first_name" json:"first_name"`
	LastName  string             `bson:"last_name" json:"last_name"`
	Gender    Gender             `bson:"gender" json:"gender"`
	Email     string             `bson:"email" json:"email"`
	Address   Address            `bson:"address" json:"address"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

type Address struct {
	Province    string `bson:"province" json:"province"`
	District    string `bson:"district" json:"district"`
	SubDistrict string `bson:"sub_district" json:"sub_district"`
	Postcode    string `bson:"postcode" json:"postcode"`
}

const (
	CreateUserError  = "SUP-USR-40001"
	GetUserError     = "SUP-USR-40002"
	GetUserListError = "SUP-USR-40003"
	UpdateUserError  = "SUP-USR-40004"
	DeleteUserError  = "SUP-USR-40005"
)
