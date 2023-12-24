package main

import (
	"database/sql"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID    `json:"id"`
	Name      string       `json:"name"`
	APIKey    string       `json:"api_key"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

type Feed struct {
	ID        uuid.UUID     `json:"id"`
	UserId    uuid.NullUUID `json:"user_id"`
	Name      string        `json:"name"`
	Url       string        `json:"url"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

func dbFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		UserId:    dbFeed.UserID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	var feeds []Feed

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, dbFeedToFeed(dbFeed))
	}

	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID    `json:"id"`
	UserId    uuid.UUID    `json:"user_id"`
	FeedId    uuid.UUID    `json:"feed_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

func dbFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		UserId:    dbFeedFollow.UserID,
		FeedId:    dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func dbFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	var feedFollows []FeedFollow

	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, dbFeedFollowToFeedFollow(dbFeedFollow))
	}

	return feedFollows
}

type Post struct {
	ID          uuid.UUID  `json:"id"`
	FeedID      uuid.UUID  `json:"feed_id"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt time.Time  `json:"published_at"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func dbPostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	var createdAt, updatedAt *time.Time
	if dbPost.CreatedAt.Valid {
		createdAt = &dbPost.CreatedAt.Time
	}

	if dbPost.UpdatedAt.Valid {
		updatedAt = &dbPost.UpdatedAt.Time
	}

	return Post{
		ID:          dbPost.ID,
		FeedID:      dbPost.FeedID,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func dbPostsToPosts(dbPosts []database.Post) []Post {
	var posts []Post

	for _, dbPost := range dbPosts {
		posts = append(posts, dbPostToPost(dbPost))
	}

	return posts
}
