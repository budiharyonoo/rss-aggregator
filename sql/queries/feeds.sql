-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT *
FROM feeds;

-- name: GetNextFetchFeeds :many
SELECT *
FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT $1;

-- name: MarkFeedAsFetch :one
UPDATE feeds
SET last_fetched_at = now(),
    updated_at      = now()
WHERE id = $1
RETURNING *;