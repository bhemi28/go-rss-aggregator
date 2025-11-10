package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bhemi28/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

// now this acts as a handler for the /user endpoint, and als has the access to the database
func (a *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	user, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJson(w, 200, convertUser(user))
}

func (a *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, convertUser(user))
}
