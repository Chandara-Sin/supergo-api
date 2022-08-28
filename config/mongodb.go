package config

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	Client *mongo.Client
}

func InitMongoDB(ctx context.Context) *DB {

	credential := options.Credential{
		Username: viper.GetString("mongo.user"),
		Password: viper.GetString("mongo.password"),
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("mongo.uri")).SetAuth(credential))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	conn := &DB{
		Client: client,
	}

	return conn

}
