package main

import (
	"net/http"

	"github.com/bhemi28/go-rss-aggregator/internal/auth"
	"github.com/bhemi28/go-rss-aggregator/internal/database"
)

type authMiddleware func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuthHandler(handler authMiddleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 401, "Unauthorized")
			return
		}

		user, err := cfg.DB.GetUserByKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, "User not found")
			return
		}

		handler(w, r, user)

	}
}
