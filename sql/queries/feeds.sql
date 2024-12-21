-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds
ORDER BY created_at DESC;

-- name: GetFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at DESC NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetched :one
UPDATE feeds
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING *;
