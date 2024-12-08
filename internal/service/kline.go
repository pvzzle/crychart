package service

import (
	"github.com/pvzzle/crychart/internal/entity"
)

type KlineProvider interface {
	GetKlines(options entity.KlineOptions) (entity.Klines, error)
}

type KlineRepository interface {
	Insert(entity.Klines) error
}

type KlineService struct {
	provider   KlineProvider
	repository KlineRepository
}

func NewKlineService(klineProvider KlineProvider, klineRepository KlineRepository) *KlineService {
	return &KlineService{
		provider:   klineProvider,
		repository: klineRepository,
	}
}

func (ks *KlineService) Fetch(options entity.KlineOptions) entity.Klines {
	klines, err := ks.provider.GetKlines(options)
	if err != nil {
		return entity.Klines{}
	}
	return klines
}

func (ks *KlineService) Save(kls entity.Klines) {
	ks.repository.Insert(kls)
}
