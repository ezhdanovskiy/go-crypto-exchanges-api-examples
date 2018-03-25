package cry

import "fmt"

type ExchangeMarketOrdersRequest struct {
	method string
	name   string
	url    string
}

func GetMarketOrders() {
	requests := []ExchangeMarketOrdersRequest{
		{
			method: "GET",
			name:   "Bit-Z",
			url:    "https://www.bit-z.com/api_v1/depth?coin=otn_btc",
		},
		{
			method: "GET",
			name:   "Cryptopia",
			url:    "https://www.cryptopia.co.nz/api/GetMarketOrders/OTN_BTC/3",
		},
		{
			method: "GET",
			name:   "HitBTC",
			url:    "https://api.hitbtc.com/api/2/public/orderbook/OTNBTC?limit=3",
		},
		{
			method: "GET",
			name:   "LiveCoin",
			url:    "https://api.livecoin.net/exchange/order_book?currencyPair=OTN/BTC&depth=3",
		},
		{
			method: "GET",
			name:   "YoBit",
			url:    "https://yobit.net/api/3/depth/otn_btc?limit=3",
		},
	}
	for _, request := range requests {
		result, err := DoRequest(request.method, request.url, nil, nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s:\n%s\n", request.name, result)
	}
}
