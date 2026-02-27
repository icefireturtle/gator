package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBUrl string `json:"db_url"`
	User  string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home, configFileName)
	return path, nil
}

func Read() (Config, error) {

	file, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	var configFile Config
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Error opening config file: %s", err)
		return Config{}, err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&configFile)
	if err != nil {
		fmt.Printf("Error decoding parameters: %s", err)
		return Config{}, err
	}
	return configFile, nil
}

func write(data *Config) (*Config, error) {

	file, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	marsh, err := json.Marshal(*data)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(file, marsh, 0644)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Config) SetUser(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("user name cannot be empty")
	}

	c.User = name

	_, err := write(c)
	if err != nil {
		return err
	}
	return nil
}
