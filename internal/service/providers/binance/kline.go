package binance

import (
	"fmt"

	"github.com/pvzzle/crychart/internal/entity"
)

func (bp *BinanceProvider) GetKlines(options entity.KlineOptions) (entity.Klines, error) {
	rawKlines, err := bp.connector.GetKlines(options.Symbol, options.Interval, options.Limit)
	if err != nil {
		fmt.Errorf("provider: failed to fetch klines: %s", err)
		return nil, err
	}

	normalizedKlines := entity.Klines{}
	for _, val := range rawKlines {
		normalizedKline := entity.Kline{
			OpenTime:  val.OpenTime,
			Open:      val.Open,
			High:      val.High,
			Low:       val.Low,
			Close:     val.Close,
			Volume:    val.Volume,
			CloseTime: val.CloseTime,
		}

		normalizedKlines = append(normalizedKlines, normalizedKline)
	}

	return normalizedKlines, nil
}
