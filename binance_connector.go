package main

type BinanceConnector struct {
	BaseURL string
}

func NewBinanceConnector() *BinanceConnector {
	return &BinanceConnector{
		BaseURL: "https://api.binance.com",
	}
}
