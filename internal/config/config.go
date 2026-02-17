package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	User  string `json:"user"`
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

func Write(data Config) Config {

	cfg := Read(configFileName)

	for key, value := range map[string]string{
		"db_url":            cfg.DBUrl,
		"current_user_name": cfg.User,
	} {
		if strings.TrimSpace(value) == "" {
			switch key {
			case "db_url":
				cfg.DBUrl = data.DBUrl
			case "current_user_name":
				cfg.User = data.User
			}
		}
	}

	err := os.WriteFile(configFileName, []byte(fmt.Sprintf(`{"db_url": "%s", "current_user_name": "%s"}`, cfg.DBUrl, cfg.User)), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s with error %s", configFileName, err)
	}
	return cfg
}

func SetUser(username string) {
	if len(username) == 0 {
		username = "username"
	}

	newUser := username

	Write(Config{User: newUser})

}
