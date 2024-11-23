package main

import (
	"fmt"
)

func main() {
	connector := NewBinanceConnector()
	resp, err := connector.Time()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(resp))
}
