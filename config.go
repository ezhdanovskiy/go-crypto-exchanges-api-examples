package cry

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	BitZ struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	} `json:"bitz"`
	HitBTC struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	} `json:"hitbtc"`
	LiveCoin struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	} `json:"livecoin"`
	YoBit struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	} `json:"yobit"`
	Cryptopia struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
	} `json:"cryptopia"`
}

func LoadConfig(fileName string) (*Config, error) {
	configFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)

	config := &Config{}
	err = jsonParser.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func SaveConfig(fileName string, config Config) error {
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, bytes, 0644)
}
