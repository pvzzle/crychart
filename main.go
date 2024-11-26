package main

import (
	"fmt"

	"github.com/pvzzle/crychart/db"
)

func main() {
	connector := NewBinanceConnector()
	_, err := connector.GetKlines("ETHBTC", "5m", 5)
	if err != nil {
		fmt.Println(err)
	}

	mongoUri := "mongodb://localhost:27017"
	db.ConnectDB(mongoUri)
}
