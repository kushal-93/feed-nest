package main

import (
	"context"
	"log"
	"time"
	"sync"

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
			log.Println("Found post: ", item.Title)
		}

		log.Printf("Feed %s collected. %v posts found.\n", feed.Name, len(rssFeed.Channel.Item))

		_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	
		if err != nil {
			log.Println("error marking feed as fetched: ", err)
		}

	}

}
