package services

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/databases"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/util"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/categories"
	"google.golang.org/grpc/grpclog"
	"log"
)

type CategoriesService struct {
	db *databases.Categories
	categories.UnimplementedCategoriesServer
	log grpclog.LoggerV2
}

func (c *CategoriesService) AddCategory(context context.Context,
	req *categories.AddCategoryRequest) (*categories.AddCategoryReply, error) {
	log.Printf("Received an AddCategory request")
	insertedId, err := c.db.AddCategory(context, req)
	if err != nil {
		c.log.Fatalln("Error inserting category: %v", err)
	}
	return &categories.AddCategoryReply{Id: insertedId}, nil
}

func NewCategoriesService() *CategoriesService {
	db := databases.NewCategories()

	return &CategoriesService{
		db:  db,
		log: util.GetGrpcLoggerV2(),
	}
}
