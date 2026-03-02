-- name: AddFeed :one

INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedsWithUsername :many
SELECT f.name, f.url, u.name
FROM feeds f
JOIN users u on f.user_id = u.id;