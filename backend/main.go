package main

import (
	"fmt"
	"github.com/Erickype/StoreProductsRecommenderBackend/protogen/golang/categories"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
)

func main() {
	category := categories.Category{
		Id:     2,
		Parent: 1,
		Name:   "Yamadory",
	}
	bytes, err := protojson.Marshal(&category)
	if err != nil {
		log.Fatalln("deserialization error", err.Error())
	}
	fmt.Println(string(bytes))
}
