package services

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/products"
	"log"
	"math/rand/v2"
)

type ProductsService struct {
	products.UnimplementedProductsServer
}

func NewProductsService() ProductsService {
	return ProductsService{}
}

func (p *ProductsService) AddProduct(_ context.Context,
	req *products.AddProductRequest) (*products.AddProductReply, error) {
	log.Printf("Received an AddProduct request")
	log.Printf("Body %v:", req.Characteristics)
	createdId := rand.Int64N(100)
	return &products.AddProductReply{Id: createdId}, nil
}
