package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request %w", err)
	}

	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request %w", err)
	}
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("invalid response %s", res.Status)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error in response %w", err)
	}

	var feed RSSFeed

	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling XML %v", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}

func aggHandler(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: agg <time between requests (e.g. 1m/1s/1h)>")
	}

	time_between_requests, _ := time.ParseDuration(cmd.args[0])
	fmt.Printf("Collecting feeds every %v\n", time_between_requests)

	ticker := time.NewTicker(time_between_requests)

	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feeds: %w\n", err)
		}
	}
}

func addFeedHandler(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	fparams := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.AddFeed(context.Background(), fparams)
	if err != nil {
		return fmt.Errorf("error adding feed %w", err)
	}

	wparams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	follows, err := s.db.CreateFeedFollow(context.Background(), wparams)
	if err != nil {
		return fmt.Errorf("error creating feed follow %w", err)
	}

	fmt.Println("Successfully added feed to database")
	fmt.Printf("* Name: %s\n", feed.Name)
	fmt.Printf("* Url: %s\n", feed.Url)
	fmt.Printf("* ID: %s\n", feed.ID)
	fmt.Printf("* Created At: %s\n", feed.CreatedAt)
	fmt.Printf("* Updated At: %s\n", feed.UpdatedAt)
	fmt.Printf("* User ID: %s\n", feed.UserID)
	fmt.Printf("* Follow ID: %s\n", follows.ID)

	return nil
}

func scrapeFeeds(s *state) error {
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), next.ID)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched %w", err)
	}

	fetched, err := fetchFeed(context.Background(), next.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed %w", err)
	}

	for _, item := range fetched.Channel.Item {
		fmt.Printf("* Title: %s\n", item.Title)
	}

	return nil
}
