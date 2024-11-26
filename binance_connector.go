package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type BinanceConnector struct {
	BaseURL string
}

func NewBinanceConnector() *BinanceConnector {
	return &BinanceConnector{
		BaseURL: "https://api.binance.com",
	}
}

func (bc *BinanceConnector) doRequest(endpoint string, params url.Values) ([]byte, error) {
	queryString := params.Encode()
	fullURL := bc.BaseURL + endpoint + "?" + queryString

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (bc *BinanceConnector) Ping() ([]byte, error) {
	return bc.doRequest("/api/v3/ping", url.Values{})
}

func (bc *BinanceConnector) Time() ([]byte, error) {
	return bc.doRequest("/api/v3/time", url.Values{})
}

func (bc *BinanceConnector) GetKlines(symbol, interval string, limit int) (Klines, error) {
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
