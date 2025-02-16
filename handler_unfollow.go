package main

import (
	"context"
	"fmt"

	"github.com/RomainDussuchal/go_projects/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>", cmd.Name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	err = s.db.DeleteFeed(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Println("Feed unfollowed successfully.")
	return nil
}