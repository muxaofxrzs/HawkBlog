package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ClientMo *mongo.Client

func NewMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://root:hawk123@1.94.27.198:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println(err)
	}
	ClientMo = client
}
func CloseMongo() {
	_ = ClientMo.Disconnect(context.Background())
}
