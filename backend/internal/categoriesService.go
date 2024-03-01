package internal

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/categories"
	"log"
)

type CategoriesService struct {
	db *DB
	categories.UnimplementedCategoriesServer
}

func NewCategoriesService(db *DB) CategoriesService {
	return CategoriesService{db: db}
}

func (c *CategoriesService) PostCategory(_ context.Context,
	req *categories.PostRequest) (*categories.PostResponse, error) {
	log.Printf("Received a PostCategory request")
	createdId, err := c.db.AddCategory(req.GetParent(), req.GetName())
	return &categories.PostResponse{Id: createdId}, err
}
