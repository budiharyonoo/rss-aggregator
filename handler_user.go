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

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	// Decode json payload
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	userPayload := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: sql.NullTime{Time: time.Now().Local(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now().Local(), Valid: true},
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), userPayload)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't create user: %s", err))
		return
	}

	respondWithJSON(w, 201, dbUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, dbUserToUser(user))
}
