package main

import (
	"context"
	"fmt"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/categories"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/products"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	// Set up a connection to the order server.
	grpcServerAddress := "localhost:50051"
	conn, err := grpc.Dial(grpcServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to GRPC server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("error closing the connection: %v", err)
		}
	}(conn)
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	if err = categories.RegisterCategoriesHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the Categories server: %v", err)
	}
	if err = products.RegisterProductsHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the Products server: %v", err)
	}
	// start listening to requests from the gateway server
	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
