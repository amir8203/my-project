-- name: CreateUser :one
INSERT INTO users (
  username, name, phone, password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;


-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1 LIMIT 1;


-- name: UpdateUsername :exec
UPDATE users
  set username = $2
WHERE id = $1;


-- name: UpdatePassword :exec
UPDATE users
  set password = $2
WHERE id = $1;


-- name: UpdateUserName :exec
UPDATE users
  set name = $2
WHERE id = $1;


-- name: UpdateUserPhone :exec
UPDATE users
  set phone = $2
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;