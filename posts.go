package main

import (
	"context"
	"fmt"
	"strconv"

	"gator/internal/database"
)

func browseHandler(s *state, cmd command, user database.User) error {

	var limit int32

	if len(cmd.args) != 1 {
		fmt.Printf("usage: browse <number of posts> (e.g. 1)\n")
		fmt.Printf("invalid argument provided, defaulting to 2\n")
		limit = 2
	} else {
		converted, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			fmt.Printf("invalid argument provided defaulting to 2\n")
			limit = 2
		}
		limit = int32(converted)
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	if len(posts) == 0 {
		fmt.Printf("no posts to browse\n")
		return nil
	}

	for _, post := range posts {
		fmt.Printf("* Title: %s\n", post.Title.String)
		fmt.Printf("* URL: %s\n", post.Url.String)
		fmt.Printf("* Description: %s\n", post.Description.String)
		fmt.Printf("* Published At: %s\n", post.PublishedAt)
		fmt.Printf("* Updated At: %s\n", post.UpdatedAt)
		fmt.Printf("* Created At: %s\n", post.CreatedAt)
		fmt.Printf("* Feed ID: %s\n", post.FeedID)
		fmt.Printf("* Post ID: %s\n", post.ID)
		fmt.Printf("\n")
	}
	return nil
}
