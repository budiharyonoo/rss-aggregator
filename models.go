package main

import (
	"database/sql"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"github.com/google/uuid"
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
