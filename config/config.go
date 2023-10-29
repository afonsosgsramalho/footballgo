package config

import (
	"io/ioutil"
	"encoding/json"
)

const APIEndpoint = "https://api.football-data.org/v4/"

type Config struct {
    APIToken string `json:"api_token"`
}

func LoadConfig(filename string) (Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return Config{}, err
    }

    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return Config{}, err
    }

    return config, nil
}