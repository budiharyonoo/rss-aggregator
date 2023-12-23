package main

import (
	"fmt"
	"github.com/budiharyonoo/rss-aggregator/internal/auth"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(&r.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Unauthorized: %s", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("User not found: %s", apiKey))
			return
		}

		handler(w, r, user)
	}
}
