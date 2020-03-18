-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users DEFAULT VALUES
RETURNING *;

-- name: GetPlayer :one
SELECT * FROM players
WHERE id = $1;

-- name: GetPlayerBySteamID :one
SELECT * FROM players
WHERE steam_id = $1;

-- name: CreatePlayer :one
INSERT INTO players (steam_id, user_id) VALUES($1, $2)
RETURNING *;

