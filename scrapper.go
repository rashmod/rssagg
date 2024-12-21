package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/rashmod/rssagg/internal/database"
)

func scrapper(db *database.Queries, concurrency int, timeOut time.Duration) {
	log.Printf("Scrapper started with concurrency %d and timeout %v", concurrency, timeOut)

	ticker := time.NewTicker(timeOut)
	for ; ; <-ticker.C {
		feedsToFetch, err := db.GetFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error getting feeds to fetch:", err)
			continue
		}

		waitGroup := &sync.WaitGroup{}

		for _, feed := range feedsToFetch {
			waitGroup.Add(1)
			go fetchFeed(waitGroup, db, feed)
		}

		waitGroup.Wait()

	}
}

func fetchFeed(waitGroup *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer waitGroup.Done()

	log.Printf("Fetching feed %s", feed.Url)

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed %s as fetched: %v", feed.Url, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s: %v", feed.Url, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.PubDate
			description.Valid = true
		}

		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %s: %v", item.PubDate, err)
			continue
		}

		_, err = db.CreatePosts(context.Background(), database.CreatePostsParams{
			ID:          uuid.New(),
			FeedID:      feed.ID,
			Title:       item.Title,
			Description: description,
			Url:         item.Link,
			PublishedAt: pubDate,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error creating item %s: %v", item.Title, err)
			continue
		}
	}

	log.Printf("Feed %s has %d items", feed.Url, len(rssFeed.Channel.Item))
}
