-- name: AddUserFeedLink :one
INSERT INTO user_feed_link (user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;
