-- name: CreateAirline :one
INSERT INTO
    airlines (
        iata,
        name,
        is_active
    )
VALUES
    (?, ?, ?)
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
