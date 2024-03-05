package services

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/databases"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/categories"
	"log"
)

type CategoriesService struct {
	db *databases.DB
	categories.UnimplementedCategoriesServer
}

func NewCategoriesService(db *databases.DB) CategoriesService {
	return CategoriesService{db: db}
}

func (c *CategoriesService) PostCategory(_ context.Context,
	req *categories.PostRequest) (*categories.PostResponse, error) {
	log.Printf("Received a PostCategory request")
	createdId, err := c.db.AddCategory(req.GetParent(), req.GetName())
	return &categories.PostResponse{Id: createdId}, err
}
