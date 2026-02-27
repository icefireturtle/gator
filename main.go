package main

import (
	"fmt"

	"gator/internal/config"
)

func main() {
	fmt.Println("Welcome to Gator!")
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config file: %s", err)
		return
	}

	cfg.SetUser("lane")

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading config file: %s", err)
		return
	}

	fmt.Printf("Database URL: %s\n", cfg.DBUrl)
	fmt.Printf("Set user %s in config file\n", cfg.User)

}
