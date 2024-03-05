// ./cmd/server/main.go
package main

import (
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/databases"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/services"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/categories"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/products"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	const addr = "0.0.0.0:50051"
	// create a TCP listener on the specified port
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a gRPC server instance
	server := grpc.NewServer()
	// create an order service instance with a reference to the db
	db := databases.NewDB()
	orderService := services.NewCategoriesService(db)
	// register the order service with the grpc server
	categories.RegisterCategoriesServer(server, &orderService)
	productsService := services.NewProductsService()
	products.RegisterProductsServer(server, &productsService)
	// start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
