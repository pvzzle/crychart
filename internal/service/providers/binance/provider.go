package binance

import "github.com/pvzzle/crychart/tools/connectors/binance"

type BinanceProvider struct {
	connector *binance.Connector
}

func NewDefaultProvider() *BinanceProvider {
	connector := binance.NewConnector()
	return NewProvider(connector)
}

func NewProvider(connector *binance.Connector) *BinanceProvider {
	return &BinanceProvider{
		connector: connector,
	}
}
