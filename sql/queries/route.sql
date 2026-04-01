-- name: CreateRoute :one
INSERT INTO
    routes (
        flightsfrom_id,
        airline_id,
        airport_from_id,
        airport_to_id,
        is_active
    )
VALUES
    (?, ?, ?, ?, ?)
RETURNING
    *;

-- name: GetRouteByFlightsFromId :one
SELECT
    *
FROM
    routes
WHERE
    routes.flightsfrom_id = ?
LIMIT
    1;

-- name: GetRoute :one
SELECT
    *
FROM
    routes
WHERE
    routes.id = ?
LIMIT
    1;
