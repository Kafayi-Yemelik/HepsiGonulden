package Customer

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(col *mongo.Collection) *Repository {
	return &Repository{
		collection: col,
	}
}
