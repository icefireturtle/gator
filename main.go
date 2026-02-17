package main

import (
	"fmt"

	"gator/internal/config"
)

func main() {
	fmt.Println("Welcome to Gator!")
	cfg := config.Read(".gatorconfig.json")
	fmt.Printf("Database URL: %s\n", cfg.DBUrl)

	config.SetUser("lane")
	fmt.Printf("Set user %s in config file", cfg.User)

}
