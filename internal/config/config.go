package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBUrl string `json:"db_url"`
}

func Read(file string) Config {
	var configFile Config
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening config file: %s", err)
		return Config{}
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&configFile)
	if err != nil {
		fmt.Printf("Error decoding parameters: %s", err)
		return Config{}
	}
	return configFile
}
