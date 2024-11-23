package main

import "fmt"

func main() {
	connector := NewBinanceConnector()
	fmt.Println(connector.BaseURL)
}
