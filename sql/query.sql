-- name: GetMessage :one
SELECT * FROM "messages"
WHERE id = $1 LIMIT 1;

-- name: GetMessageIds :many
SELECT id FROM "messages"
ORDER BY id;