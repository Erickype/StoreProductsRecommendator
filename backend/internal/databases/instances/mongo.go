package instances

import (
	"context"
	"github.com/Erickype/StoreProductsRecommenderBackend/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/grpclog"
)

type MongoConnection struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
	Log        grpclog.LoggerV2
}

func (mc *MongoConnection) InitCollection(database, collection string) {
	mc.Database = mc.Client.Database(database)
	mc.Collection = mc.Database.Collection(collection)
}

func (mc *MongoConnection) Disconnect() {
	if err := mc.Client.Disconnect(context.TODO()); err != nil {
		mc.Log.Fatalf("Error disconnecting: %v", err)
	}
}

func NewMongoConnection() *MongoConnection {
	mc := &MongoConnection{Log: util.GetGrpcLoggerV2()}

	uri := util.MongodbUri.String()
	if uri == "" {
		mc.Log.Fatalln("You must set your 'MONGODB_URI' environment variable.")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		mc.Log.Fatalf("Error connecting: %v\n", err)
	}
	mc.Client = client
	return mc
}
