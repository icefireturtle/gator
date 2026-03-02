package main

import (
	"context"
	"fmt"
	"os"
)

func resetHandler(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to reset database table")
	}
	fmt.Println("Successfully reset users table")
	os.Exit(0)
	return nil
}
