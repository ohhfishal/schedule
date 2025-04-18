-- name: CreateEvent :one
INSERT INTO events (
  name,
  description,
  start_time,
  recurrence
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateEvent :one
UPDATE events
  SET 
    name = coalesce(sqlc.narg('name'), name),
    description = coalesce(sqlc.narg('description'), description),
    start_time = coalesce(sqlc.narg('start_time'), start_time),
    end_time = coalesce(sqlc.narg('end_time'), end_time),
    recurrence = coalesce(sqlc.narg('recurrence'), recurrence)
WHERE id = sqlc.arg('id')
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

-- name: DeleteEvent :execresult
DELETE FROM events
WHERE id = ?;
