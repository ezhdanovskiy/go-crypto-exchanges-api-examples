package cry

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetOpenOrders(config Config) {
	fmt.Printf("config: %+v", config)
	//result, err := GetOpenOrdersBitZ(config)
	result, err := GetOpenOrdersCryptopia(config)
	//result, err := GetOpenOrdersHitBTC(config)
	//result, err := GetOpenOrdersLiveCoin(config)
	//result, err := GetOpenOrdersYoBit(config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result:\n%s\n", result)
}
func GetOpenOrdersBitZ(config Config) ([]byte, error) {
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)[:6]
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	values := url.Values{
		"api_key":   []string{config.BitZ.ApiKey},
		"coin":      []string{"otn_btc"},
		"nonce":     []string{nonce},
		"timestamp": []string{timestamp},
	}

	md5Sum := md5.Sum([]byte(values.Encode() + config.BitZ.ApiSecret))
	sign := fmt.Sprintf("%x", md5Sum)
	values["sign"] = []string{sign}

	method := "POST"
	refURL := "https://www.bit-z.com/api_v1/openOrders"
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}
	return DoRequest(method, refURL, strings.NewReader(values.Encode()), headers)
}

func GetOpenOrdersCryptopia(config Config) ([]byte, error) {
	method := "POST"
	refURL := "https://www.cryptopia.co.nz/api/GetOpenOrders"
	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	fmt.Println("nonce", nonce)

	values := url.Values{
		"Market": []string{"OTN/BTC"},
	}

	md5Sum := md5.Sum([]byte(values.Encode()))
	hashedPostParams := base64.StdEncoding.EncodeToString(md5Sum[:])
	signature := config.Cryptopia.ApiKey + "POST" + strings.ToLower(url.QueryEscape(refURL)) + nonce + hashedPostParams
	key, _ := base64.StdEncoding.DecodeString(config.Cryptopia.ApiSecret)
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(signature))
	hmacSignature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	auth := "amx " + config.Cryptopia.ApiKey + ":" + hmacSignature + ":" + nonce

	headers := map[string]string{
		"Authorization": auth,
		"Content-Type":  "application/x-www-form-urlencoded",
	}
	return DoRequest(method, refURL, strings.NewReader(values.Encode()), headers)
}

func GetOpenOrdersHitBTC(config Config) ([]byte, error) {
	method := "GET"
	refURL := "https://api.hitbtc.com/api/2/order"
	headers := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString(
			[]byte((config.HitBTC.ApiKey + ":" + config.HitBTC.ApiSecret))),
	}
	return DoRequest(method, refURL, nil, headers)
}

func GetOpenOrdersLiveCoin(config Config) ([]byte, error) {
	method := "GET"
	refURL := "https://api.livecoin.net/exchange/client_orders"
	values := url.Values{
		"api_key": []string{"OPEN"},
	}

	mac := hmac.New(sha256.New, []byte(config.LiveCoin.ApiSecret))
	mac.Write([]byte(values.Encode()))
	sign := hex.EncodeToString(mac.Sum(nil))

	headers := map[string]string{
		"Api-key": config.LiveCoin.ApiKey,
		"Sign":    strings.ToUpper(sign),
	}
	return DoRequest(method, refURL, strings.NewReader(values.Encode()), headers)
}

func GetOpenOrdersYoBit(config Config) ([]byte, error) {
	method := "POST"
	refURL := "https://yobit.net/tapi"
	nonce := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println("nonce", nonce)
	values := url.Values{
		"method": []string{"ActiveOrders"},
		"nonce":  []string{nonce},
		"pair":   []string{"otn_btc"},
	}

	mac := hmac.New(sha512.New, []byte(config.YoBit.ApiSecret))
	mac.Write([]byte(values.Encode()))
	sign := hex.EncodeToString(mac.Sum(nil))

	headers := map[string]string{
		"Key":          config.YoBit.ApiKey,
		"Sign":         sign,
		"Content-Type": "application/x-www-form-urlencoded",
	}
	return DoRequest(method, refURL, strings.NewReader(values.Encode()), headers)
}
