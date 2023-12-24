-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedByUserId :many
SELECT feed_follows.*
FROM feed_follows
         LEFT JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE
FROM feed_follows
WHERE feed_id = $1
  AND user_id = $2;