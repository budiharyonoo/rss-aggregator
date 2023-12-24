package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	// Decode json payload
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feedFollowPayload := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
		CreatedAt: sql.NullTime{Time: time.Now().Local(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now().Local(), Valid: true},
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), feedFollowPayload)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't create feed: %s", err))
		return
	}

	respondWithJSON(w, 201, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedByUserId(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't get feeds: %s", err))
		return
	}
	respondWithJSON(w, 200, dbFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdParam := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdParam)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't parse Feed Follow ID: %s", err))
	}

	feedFollowPayload := database.DeleteFeedFollowParams{
		FeedID: feedFollowId,
		UserID: user.ID,
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), feedFollowPayload)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't delete feed follow: %s", err))
		return
	}
	respondWithJSON(w, 204, struct{}{})
}
