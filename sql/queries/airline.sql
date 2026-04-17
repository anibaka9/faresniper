-- name: CreateAirline :one
INSERT INTO
    airlines (
        flightsfrom_id,
        iata,
        name,
        is_active
    )
VALUES
    (?, ?, ?, ?)
RETURNING
    *;

-- name: GetAirlineByFlightsFromId :one
SELECT
    *
FROM
    airlines
WHERE
    airlines.flightsfrom_id = ?
LIMIT
    1;

-- name: GetAirline :one
SELECT
    *
FROM
    airlines
WHERE
    airlines.id = ?
LIMIT
    1;

-- name: GetAirlines :many
SELECT
    *
FROM
    airlines
LIMIT
    ? OFFSET ?;

-- name: CountAirlines :one
SELECT
    COUNT(*) AS cont_airlines
FROM
    airlines;

-- name: UpdateAirline :one
UPDATE
    airlines
SET
    flightsfrom_id = ?,
    iata = ?,
    name = ?,
    is_active = ?
WHERE
    id = ?
RETURNING
    *;
