package main

import (
	"context"
	"fmt"
	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

func feedFollowsHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retreiving feed: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.User)
	if err != nil {
		return fmt.Errorf("user does not exist: %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed,
		UserID:    user.ID,
	}

	follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error creating follow %w", err)
	}

	fmt.Printf("Feed %s has been successfully added for user %s", follow.FeedName, follow.UserName)

	return nil
}
