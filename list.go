package main

import (
	"context"
	"fmt"
)

func listUserHandler(s *state, cmd command) error {
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

func listFeedHandler(s *state, cmd command) error {
	feeds, err := s.db.GetFeedsWithUsername(context.Background())
	if err != nil {
		return fmt.Errorf("unable to retreive feed data")
	}

	for _, feed := range feeds {
		fmt.Printf("* Name: %s\n", feed.Name)
		fmt.Printf("* Url: %s\n", feed.Url)
		fmt.Printf("* User: %s\n", feed.Name_2)
	}
	return nil
}
