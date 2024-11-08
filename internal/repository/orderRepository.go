package repository

import (
	"HepsiGonulden/internal/types"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(client *mongo.Client) (*OrderRepository, error) {

	dbName := viper.GetString("database.order.dbName")
	collectionName := viper.GetString("database.order.collectionName")

	return &OrderRepository{collection: client.Database(dbName).Collection(collectionName)}, nil
}

func (r *OrderRepository) FindByID(ctx context.Context, id string) (*types.Order, error) {
	var order *types.Order

	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no order found with ID %s", id)
		}
	}
	return order, nil
}
func (r *OrderRepository) OrderCreate(ctx context.Context, order *types.Order) (*mongo.InsertOneResult, error) {
	res, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *OrderRepository) OrderUpdate(ctx context.Context, id string, order *types.Order) error {
	filter := bson.D{{"_id", id}}
	update := bson.M{"$set": order}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *OrderRepository) OrderDelete(ctx context.Context, id string) error {
	filter := bson.D{{"_id", id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
