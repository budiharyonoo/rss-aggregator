-- name: CreatePost :one
INSERT INTO posts (id, feed_id, title, url, description, published_at, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostByUserId :many
SELECT p.*
FROM posts p
         JOIN feed_follows ff on p.feed_id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY p.published_at DESC
LIMIT $2;