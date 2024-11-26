package main

import (
	"encoding/json"
	"fmt"
)

type Kline struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	TradeNum                 int64  `json:"tradeNum"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

type Klines []Kline

func parseKline(raw []json.RawMessage, kline *Kline) error {
	fields := []interface{}{
		&kline.OpenTime,
		&kline.Open,
		&kline.High,
		&kline.Low,
		&kline.Close,
		&kline.Volume,
		&kline.CloseTime,
		&kline.QuoteAssetVolume,
		&kline.TradeNum,
		&kline.TakerBuyBaseAssetVolume,
		&kline.TakerBuyQuoteAssetVolume,
	}

	if len(raw) < len(fields) {
		return fmt.Errorf("unexpected kline format: %v", raw)
	}

	for i, field := range fields {
		if err := json.Unmarshal(raw[i], field); err != nil {
			return fmt.Errorf("error parsing field %d: %w", i, err)
		}
	}

	return nil
}
