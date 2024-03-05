package main

import (
	"github.com/Erickype/StoreProductsRecommenderBackend/gateway"
	"github.com/Erickype/StoreProductsRecommenderBackend/insecure"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/categories"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	server := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)
	db := internal.NewDB()
	orderService := internal.NewCategoriesService(db)
	// register the order service with the grpc server
	categories.RegisterCategoriesServer(server, &orderService)
	productsService := internal.NewProductsService()
	products.RegisterProductsServer(server, &productsService)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(server.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
