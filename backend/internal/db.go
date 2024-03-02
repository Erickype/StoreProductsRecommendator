package internal

import (
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/categories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand/v2"
)

type DB struct {
	collection []*categories.Category
}

// NewDB creates a new array to mimic the behaviour of an in-memory database
func NewDB() *DB {
	return &DB{
		collection: make([]*categories.Category, 0),
	}
}

// AddCategory adds a new category to the DB collection. Returns an error on duplicate ids
func (d *DB) AddCategory(parent int64, name string) (int64, error) {
	id := rand.Int64N(100)
	for _, o := range d.collection {
		if o.Id == id {
			return -1, status.Errorf(codes.AlreadyExists, "duplicate category id: %d", id)
		}
	}
	d.collection = append(d.collection, &categories.Category{
		Id:     id,
		Parent: parent,
		Name:   name,
	})
	return id, nil
}
