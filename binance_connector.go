package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
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

	body, err := ioutil.ReadAll(resp.Body)
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
