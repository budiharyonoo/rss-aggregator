package main

import (
	"context"
	"database/sql"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"github.com/google/uuid"
	"log"
	"strings"
	"sync"
	"time"
)

func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Println("Scrapping on", concurrency, "every", timeBetweenRequest, "duration")

	ticker := time.NewTicker(timeBetweenRequest)
	// This will run the script immediately
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFetchFeeds(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error fetching feeds:", err)
			continue // go to next unfetched feed
		}

		// We create mechanism to fetch the feed asynchrously using go routine
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1) // Spawn 1 per feed

			go scrapeFeed(wg, db, feed)
		}
		wg.Wait() // Wait until each spawned finish then go to the next loop of ticker.C
	}
}

func scrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	// Mark/flag as done fetch
	_, err := db.MarkFeedAsFetch(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error mark last fetch feed ID of:", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error fetch RSS feed from ID of:", feed.ID, "and URL of: ", feed.Url, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		// Description empty string logic
		description := sql.NullString{}
		if item.Description != "" {
			description.Valid = true
			description.String = item.Description
		}

		// Parse published_at
		publishedAt, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			log.Println("Error parsing published at RSS feed of:", item.Title, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			FeedID:      feed.ID,
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedAt,
			CreatedAt:   sql.NullTime{Time: time.Now().Local(), Valid: true},
			UpdatedAt:   sql.NullTime{Time: time.Now().Local(), Valid: true},
		})
		if err != nil {
			//// If there is error from DB, about duplicated URL
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}

			log.Println("Error saving RSS feed of:", item.Title, err)
			continue
		}
	}
	log.Println("Feed", feed.Name, "collected.", len(rssFeed.Channel.Item), "posts found.")
}
