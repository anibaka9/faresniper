-- name: CreateAirport :one
INSERT INTO
    airports (iata, name, city_id)
VALUES
    (?, ?, ?)
RETURNING
    *;

-- name: GetAirportByIata :one
SELECT
    *
FROM
    airports
WHERE
    airports.iata = ?
LIMIT
    1;
