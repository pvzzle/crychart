package binance

import (
	"io"
	"net/http"
	"net/url"
)

// API versions?
// Tests?

// TODO: add client & options
type Connector struct {
	BaseURL string
}

// TODO: add constructor with options
func NewConnector() *Connector {
	return &Connector{
		BaseURL: "https://api.binance.com",
	}
}

// TODO: rewrite & expand this terrible func
func (bc *Connector) doRequest(endpoint string, params url.Values) ([]byte, error) {
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
