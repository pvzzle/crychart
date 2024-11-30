package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Errorf("failed to connect mongodb: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		fmt.Errorf("failed to ping mongodb: %v", err)
	}

	fmt.Println("connected to mongodb")
	Client = client
}
