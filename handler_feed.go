package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	// Decode json payload
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feedPayload := database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    uuid.NullUUID{Valid: true, UUID: user.ID},
		Name:      params.Name,
		Url:       params.URL,
		CreatedAt: sql.NullTime{Time: time.Now().Local(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now().Local(), Valid: true},
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), feedPayload)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't create feed: %s", err))
		return
	}

	respondWithJSON(w, 201, dbFeedToFeed(feed))
}
