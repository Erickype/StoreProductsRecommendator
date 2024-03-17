package services

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/databases"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/util"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/categories"
	"google.golang.org/grpc/grpclog"
)

type CategoriesService struct {
	db *databases.Categories
	categories.UnimplementedCategoriesServer
	log grpclog.LoggerV2
}

func (c *CategoriesService) AddCategory(context context.Context,
	req *categories.AddCategoryRequest) (*categories.AddCategoryReply, error) {
	c.log.Infoln("Received an AddCategory request")
	insertedId, err := c.db.AddCategory(context, req)
	if err != nil {
		c.log.Errorf("Error inserting category: %v", err.Error())
		return nil, err
	}
	return &categories.AddCategoryReply{Id: insertedId}, nil
}

func (c *CategoriesService) GetCategoryById(context context.Context,
	req *categories.GetCategoryByIdRequest) (*categories.GetCategoryByIdResponse, error) {
	c.log.Infoln("Received a GetCategoryById request")
	category, err := c.db.GetCategoryById(context, req)
	if err != nil {
		c.log.Errorf("Error getting a category: %v", err.Error())
		return nil, err
	}
	return category, nil
}

func NewCategoriesService() *CategoriesService {
	db := databases.NewCategories()

	return &CategoriesService{
		db:  db,
		log: util.GetGrpcLoggerV2(),
	}
}
