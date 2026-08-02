package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kushal-93/feed-nest/internal/database"
)

func startScraping(db *database.Queries, concurrency int, interval time.Duration) {
	log.Printf("Scraping on %v goroutines every %s interval", concurrency, interval)
	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrape(db, feed, wg)
		}
		wg.Wait()

	}
}

func scrape(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("error fetching feed: ", err)
	} else {
		for _, item := range rssFeed.Channel.Item {
			description := sql.NullString{}
			if item.Description != "" {
				description.String = item.Description
				description.Valid = true
			}

			pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				log.Printf("couldn't parse date %v with error %v\n", item.PubDate, err)
			}

			_, err = db.CreatePost(context.Background(), database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
				Title:       item.Title,
				Description: description,
				PublishedAt: pubAt,
				Url:         item.Link,
				FeedID:      feed.ID,
			})

			if err != nil {
				if !strings.Contains(err.Error(), "duplicate key") {
					log.Println("failed to create post: ", err)
				}
			}
		}

		log.Printf("Feed %s collected. %v posts found.\n", feed.Name, len(rssFeed.Channel.Item))

		_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

		if err != nil {
			log.Println("error marking feed as fetched: ", err)
		}

	}

}
