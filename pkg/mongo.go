package pkg

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func GetMongoClient(timeout time.Duration, URI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	connectionString := viper.GetString("database.connection_string")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetMongoCollection(client *mongo.Client, dbName, colName string) (*mongo.Collection, error) {
	col := client.Database(dbName).Collection(colName)
	return col, nil
}
