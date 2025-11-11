package main

import (
	"time"

	"github.com/bhemi28/go-rss-aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func convertUser(user database.User) User {
	return User{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	UserID    uuid.UUID `json:"user_id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convertFeed(feed database.Feed) Feed {
	return Feed{
		UserID:    feed.UserID,
		Title:     feed.Title,
		Url:       feed.Url,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
}


type FeedFollow struct {
	FeedId    uuid.UUID `json:"feedId"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
}

func convertFeedFollow(feed database.GetUserFeedFollowsRow) FeedFollow {
	return FeedFollow {
		FeedId: feed.FeedID,
		CreatedAt: feed.CreatedAt,
		Title: feed.Title.String,
		Url: feed.Url.String,
	}
}

func convertFeedFollows (dbFeeds []database.GetUserFeedFollowsRow) []FeedFollow {
	feeds := []FeedFollow{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, convertFeedFollow(dbFeed))
	}

	return feeds
}
