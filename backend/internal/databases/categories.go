package databases

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/databases/instances"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/util"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/categories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	*instances.MongoConnection
	database   string
	collection string
}

func (c *Categories) AddCategory(context context.Context, req *categories.AddCategoryRequest) (string, error) {
	result, err := c.Collection.InsertOne(context, req)
	if err != nil {
		return "", err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func NewCategories() *Categories {
	mc := instances.NewMongoConnection()
	c := &Categories{
		MongoConnection: mc,
		database:        util.Products.String(),
		collection:      util.Categories.String(),
	}
	c.InitCollection(c.database, c.collection)
	return c
}
