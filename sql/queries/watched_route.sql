-- name: CreateWatchedRoute :one
INSERT INTO
    watched_routes(origin_iata, destination_iata, active)
VALUES
    (?, ?, ?)
RETURNING
    *;

-- name: CountWatchedRoutes :one
SELECT
    COUNT(*)
FROM
    watched_routes;

-- name: GetWatchedRoutes :many
SELECT
    *
FROM
    watched_routes
LIMIT
    ? OFFSET ?;

-- name: ChangeActive :one
UPDATE
    watched_routes
SET
    active = ?
WHERE
    origin_iata = ?
    AND destination_iata = ?
RETURNING
    *;
