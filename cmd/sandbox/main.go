package main

import (
	"github.com/pvzzle/crychart/internal/entity"
	"github.com/pvzzle/crychart/internal/service"
	"github.com/pvzzle/crychart/internal/service/providers/binance"
	"github.com/pvzzle/crychart/internal/service/repo/mongo"
)

func main() {
	provider := binance.NewDefaultProvider()

	mongoUri := "mongodb://localhost:27017"
	klinesCollection := mongo.ConnectDB(mongoUri).Database("crychart").Collection("klines")
	repo := mongo.NewKlineRepo(klinesCollection)

	klineService := service.NewKlineService(
		provider,
		repo,
	)
	klineOpts := entity.KlineOptions{
		Symbol:   "ETHBTC",
		Interval: "5m",
		Limit:    5,
	}

	kls := klineService.Fetch(klineOpts)
	klineService.Save(kls)
}
