// ./cmd/server/main.go
package main

import (
	"github.com/Erickype/StoreProductsRecommenderBackend/internal"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/categories"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/products"
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
	db := internal.NewDB()
	orderService := internal.NewCategoriesService(db)
	// register the order service with the grpc server
	categories.RegisterCategoriesServer(server, &orderService)
	productsService := internal.NewProductsService()
	products.RegisterProductsServer(server, &productsService)
	// start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
