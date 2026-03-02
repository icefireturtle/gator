package main

import (
	"context"
	"fmt"
)

func listHandler(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("unable to retreive user data")
	}

	for _, user := range users {
		if user == s.cfg.User {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}
