-- name: CreateAirport :one
INSERT INTO
    airports (iata, name, city_id)
VALUES
    (?, ?, ?)
RETURNING
    *;

-- name: GetAirport :one
SELECT
    a.id,
    a.iata,
    a.name,
    a.city_id,
    c.name AS city_name
FROM
    airports a
    JOIN cities c ON a.city_id = c.id
WHERE
    a.id = ?
LIMIT
    1;

-- name: GetAirportByIata :one
SELECT
    *
FROM
    airports
WHERE
    airports.iata = ?
LIMIT
    1;

-- name: GetAirports :many
SELECT
    a.id,
    a.iata,
    a.name,
    a.city_id,
    c.name city_name
FROM
    airports a
    JOIN cities c ON a.city_id = c.id
LIMIT
    ? OFFSET ?;

-- name: CountAirports :one
SELECT
    COUNT(*) AS count_airports
FROM
    airports;
