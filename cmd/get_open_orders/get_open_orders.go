package main

import "github.com/ezhdanovskiy/go-crypto-exchanges-api-examples"

func main() {
	config, err := cry.LoadConfig(".cry.config.json")
	if err != nil {
		panic(err)
	}
	cry.GetOpenOrders(*config)
}
