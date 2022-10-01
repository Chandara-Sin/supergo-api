package config

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client
}

func InitMongoDB(ctx context.Context) *DB {

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    viper.GetString("mongo.db"),
		Username:      viper.GetString("mongo.user"),
		Password:      viper.GetString("mongo.password"),
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("mongo.uri")).SetAuth(credential))
	if err != nil {
		panic(err)
	}

	conn := &DB{
		Client: client,
	}

	return conn

}
