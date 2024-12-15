package main

import (
	"time"

	"github.com/google/uuid"

	"github.com/rashmod/rssagg/internal/database"
)

type user struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) user {
	return user{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedToFeed(dbFeed database.Feed) feed {
	return feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []feed {
	feeds := make([]feed, len(dbFeeds))

	for i, dbFeed := range dbFeeds {
		feeds[i] = databaseFeedToFeed(dbFeed)
	}

	return feeds
}
