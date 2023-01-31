-- name: GetMessageById :one
SELECT * FROM "messages"
WHERE id = $1 LIMIT 1;

-- name: GetMessagesByIds :many
SELECT * FROM "messages"
WHERE id = ANY($1::bigint[]);

-- name: GetMessageIds :many
SELECT id FROM "messages"
ORDER BY id;