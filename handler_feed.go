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
func (a *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Title string `json:"title"`
		Url   string `json:"url"`
	}

	var params parameters
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid request payload: %v", err))
		return
	}

	feed, err := a.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Title:     params.Title,
		Url:       params.Url,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJson(w, 200, convertFeed(feed))
}

func (a *apiConfig) addFeedToUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the feed id from the url
	feedID := r.URL.Query().Get("feed_id")
	if feedID == "" {
		respondWithError(w, 400, "Missing feed_id parameter")
		return
	}

	// add the feed to the user
	_, err := a.DB.AddUserFeedLink(r.Context(), database.AddUserFeedLinkParams{
		UserID: user.ID,
		FeedID: uuid.MustParse(feedID),
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error adding feed to user: %v", err))
		return
	}

	respondWithJson(w, 200, struct{}{})
}


func (a *apiConfig) getFeedForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds, err := a.DB.GetUserFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error adding feed to user: %v", err))
		return
	}

	respondWithJson(w, 200, convertFeedFollows(feeds))
}

func (a *apiConfig) removeFeedFromUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the feed id from the url
	feedID := r.URL.Query().Get("feed_id")
	if feedID == "" {
		respondWithError(w, 400, "Missing feed_id parameter")
		return
	}

	// add the feed to the user
	_, err := a.DB.AddUserFeedLink(r.Context(), database.AddUserFeedLinkParams{
		UserID: user.ID,
		FeedID: uuid.MustParse(feedID),
	})

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error adding feed to user: %v", err))
		return
	}

	respondWithJson(w, 200, struct{}{})
}
