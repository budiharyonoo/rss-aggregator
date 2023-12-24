package main

import (
	"fmt"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"net/http"
)

func (apiCfg *apiConfig) handlerGetPostByUserId(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostByUserId(r.Context(), database.GetPostByUserIdParams{
		UserID: user.ID, Limit: int32(10),
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Can't get posts: %s", err))
	}

	respondWithJSON(w, 200, dbPostsToPosts(posts))
}
