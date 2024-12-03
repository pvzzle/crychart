package service

import (
	"github.com/pvzzle/crychart/internal/entity"
)

type KlineProvider interface {
	GetKlines(options entity.KlineOptions) (entity.Klines, error)
}

type KlineRepository interface {
	// TODO: add methods
}

type KlineService struct {
	provider   KlineProvider
	repository KlineRepository
}

func NewKlinesService(klinesProvider KlineProvider, KlinesRepository KlineRepository) *KlineService {
	return &KlineService{
		provider:   klinesProvider,
		repository: KlinesRepository,
	}
}

func (ks *KlineService) Fetch(options entity.KlineOptions) entity.Klines {
	klines, err := ks.provider.GetKlines(options)
	if err != nil {
		return entity.Klines{}
	}
	return klines
}
