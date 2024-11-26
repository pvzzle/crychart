package main

import (
	"fmt"
)

func main() {
	connector := NewBinanceConnector()
	resp, err := connector.GetKlines("ETHBTC", "5m", 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
