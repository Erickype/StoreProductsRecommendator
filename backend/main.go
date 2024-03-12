package main

import (
	"github.com/Erickype/StoreProductsRecommenderBackend/gateway"
	"github.com/Erickype/StoreProductsRecommenderBackend/insecure"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/services"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/util"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/categories"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/v1/products"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := util.GetGrpcLoggerV2()

	addr := "0.0.0.0:10000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	server := grpc.NewServer(
		// TODO: Replace with your own certificate!
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
	)

	RegisterServices(server)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(server.Serve(lis))
	}()

	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}

func RegisterServices(server *grpc.Server) {
	categoriesService := services.NewCategoriesService()
	categories.RegisterCategoriesServer(server, categoriesService)

	productsService := services.NewProductsService()
	products.RegisterProductsServer(server, &productsService)
}
