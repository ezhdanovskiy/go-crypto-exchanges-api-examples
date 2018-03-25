package main

import "github.com/ezhdanovskiy/go-crypto-exchanges-api-examples"

func main() {
	var config cry.Config
	cry.SaveConfig(".cry.config.json", config)
}
