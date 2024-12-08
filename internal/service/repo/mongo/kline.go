package mongo

import (
	"context"

	"github.com/pvzzle/crychart/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type KlineRepo struct {
	*mongo.Collection
}

func NewKlineRepo(collection *mongo.Collection) *KlineRepo {
	return &KlineRepo{
		collection,
	}
}

func (kr *KlineRepo) Insert(kls entity.Klines) error {
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

	_, err := kr.InsertMany(context.Background(), klines)
	if err != nil {
		return err
	}

	return nil
}
