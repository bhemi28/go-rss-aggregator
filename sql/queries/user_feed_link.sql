-- name: AddUserFeedLink :one
INSERT INTO user_feed_link (user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: GetUserFeedFollows :many
SELECT ufl.feed_id, ufl.created_at , f.title, f.url
FROM user_feed_link ufl
LEFT JOIN feeds f ON ufl.feed_id = f.id
WHERE ufl.user_id = $1;
