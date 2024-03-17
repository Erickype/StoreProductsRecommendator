package databases

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/databases/instances"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/util"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/categories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Categories struct {
	*instances.MongoConnection
	database   string
	collection string
}

func ProtoToBSONCategory(category *categories.Category) (bson.M, error) {
	oid, err := primitive.ObjectIDFromHex(category.GetId())
	if err != nil {
		return nil, err
	}
	return bson.M{
		"_id":    oid,
		"parent": category.GetParent(),
		"name":   category.GetName(),
	}, nil
}

func BSONToProtoCategory(data bson.M) *categories.Category {
	var category categories.Category
	category.Id = data["_id"].(primitive.ObjectID).Hex()
	category.Parent = data["parent"].(string)
	category.Name = data["name"].(string)
	return &category
}

func (c *Categories) AddCategory(context context.Context, req *categories.AddCategoryRequest) (string, error) {
	result, err := c.Collection.InsertOne(context, req)
	if err != nil {
		return "", err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

func (c *Categories) GetCategoryById(context context.Context,
	req *categories.GetCategoryByIdRequest) (*categories.GetCategoryByIdResponse, error) {
	objectID, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, err
	}
	result := c.Collection.FindOne(context, bson.D{{"_id", objectID}})
	if err != nil {
		return nil, err
	}
	var mongoCategory bson.M
	err = result.Decode(&mongoCategory)
	if err != nil {
		return nil, err
	}
	category := BSONToProtoCategory(mongoCategory)
	return &categories.GetCategoryByIdResponse{Category: category}, nil
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
