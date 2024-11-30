package main

import (
	"context"
	"fmt"

	"github.com/pvzzle/crychart/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	connector := NewBinanceConnector()
	kls, err := connector.GetKlines("ETHBTC", "5m", 5)
	if err != nil {
		fmt.Println(err)
	}

	mongoUri := "mongodb://localhost:27017"
	mongo.ConnectDB(mongoUri)

	klinesCollection := mongo.Client.Database("binance").Collection("klines")
	klines := make([]interface{}, len(kls))

	for i, kl := range kls {
		klines[i] = bson.D{
			{Key: "openTime", Value: kl.OpenTime},
			{Key: "open", Value: kl.Open},
			{Key: "high", Value: kl.High},
			{Key: "low", Value: kl.Low},
			{Key: "close", Value: kl.Close},
			{Key: "volume", Value: kl.Volume},
			{Key: "closeTime", Value: kl.CloseTime},
		}
	}

	klinesCollection.InsertMany(context.Background(), klines)
}
