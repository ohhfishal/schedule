-- name: CreateEvent :one
INSERT INTO events (
  name,
  description,
  start_time
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: UpdateEvent :one
UPDATE events
set name = ?
WHERE id = ?
RETURNING *;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = ? LIMIT 1;

-- name: GetAllEvents :many
SELECT * FROM events
ORDER BY id;

-- name: GetEvents :many
SELECT * FROM events
WHERE sqlc.arg('start') <= start_time AND start_time <= sqlc.arg('end')
ORDER BY start_time;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = ?;
