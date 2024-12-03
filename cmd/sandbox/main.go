package main

import (
	"context"

	"github.com/pvzzle/crychart/internal/entity"
	"github.com/pvzzle/crychart/internal/service"
	"github.com/pvzzle/crychart/internal/service/providers/binance"

	"github.com/pvzzle/crychart/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	provider := binance.NewDefaultProvider()

	klinesService := service.NewKlinesService(
		provider,
		nil,
	)
	klineOpts := entity.KlineOptions{
		Symbol:   "ETHBTC",
		Interval: "5m",
		Limit:    5,
	}
	kls := klinesService.Fetch(klineOpts)

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
