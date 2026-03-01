package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func registerHandler(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments provided")
	}

	if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	fmt.Printf("Created User %s in database\n", user.Name)

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("Set user %s in config file\n", user.Name)

	return nil
}
