package repository

import (
	"HepsiGonulden/internal/types"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository struct {
	collection *mongo.Collection
}

func NewCustomerRepository(client *mongo.Client) (*CustomerRepository, error) {

	dbName := viper.GetString("database.customer.dbName")
	collectionName := viper.GetString("database.customer.collectionName")

	return &CustomerRepository{collection: client.Database(dbName).Collection(collectionName)}, nil
}

func (r *CustomerRepository) FindByID(ctx context.Context, id string) (*types.Customer, error) {
	var customer *types.Customer

	filter := bson.M{"_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no customer found with ID %s", id)
		}
	}
	return customer, nil
}

func (r *CustomerRepository) FindByEmail(ctx context.Context, email string) (*types.Customer, error) {
	var customer *types.Customer
	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return customer, nil
}

func (r *CustomerRepository) Create(ctx context.Context, customer *types.Customer) (*mongo.InsertOneResult, error) {
	res, err := r.collection.InsertOne(ctx, customer)
	return res, err
}

func (r *CustomerRepository) Update(ctx context.Context, id string, customer *types.Customer) error {
	filter := bson.D{{"_id", id}}
	update := bson.M{"$set": customer}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *CustomerRepository) Delete(ctx context.Context, id string) error {
	filter := bson.D{{"_id", id}}
	_, err := r.collection.DeleteOne(ctx, filter)
	return err
}
