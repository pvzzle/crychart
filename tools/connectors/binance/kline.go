package binance

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
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

func (bc *Connector) GetKlines(symbol, interval string, limit int) (Klines, error) {
	params := url.Values{
		"symbol":   []string{symbol},
		"interval": []string{interval},
		"limit":    []string{strconv.Itoa(limit)},
	}
	data, err := bc.doRequest("/api/v3/klines", params)
	if err != nil {
		return nil, err
	}

	var rawData [][]json.RawMessage
	if err := json.Unmarshal(data, &rawData); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	klines := make(Klines, len(rawData))
	for i, raw := range rawData {
		var kline Kline
		if err := parseKline(raw, &kline); err != nil {
			return nil, fmt.Errorf("error parsing kline at index %d: %w", i, err)
		}
		klines[i] = kline
	}

	return klines, nil
}

// TODO: just do simple if-else?
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
