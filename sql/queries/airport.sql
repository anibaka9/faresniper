-- name: CreateAirport :one
INSERT INTO
    airports (iata, name, iata_type, flightable, city_iata)
VALUES
    (?, ?, ?, ?, ?)
RETURNING
    *;

-- name: GetAirport :one
SELECT
    a.iata,
    a.name,
    a.iata_type,
    a.flightable,
    a.city_iata,
    c.name AS city_name
FROM
    airports a
    JOIN cities c ON a.city_iata = c.iata
WHERE
    a.iata = ?
LIMIT
    1;

-- name: GetAirports :many
SELECT
    a.iata,
    a.name,
    a.iata_type,
    a.flightable,
    a.city_iata,
    c.name AS city_name
FROM
    airports a
    JOIN cities c ON a.city_iata = c.iata
ORDER BY
    a.iata
LIMIT
    ? OFFSET ?;

-- name: CountAirports :one
SELECT
    COUNT(*) AS count_airports
FROM
    airports;

-- name: UpdateAirport :one
UPDATE
    airports
SET
    name = ?,
    iata_type = ?,
    flightable = ?,
    city_iata = ?
WHERE
    iata = ?
RETURNING
    *;

-- name: GetAllAirports :many
SELECT
    a.iata,
    c.name AS city_name
FROM
    airports a
    JOIN cities c ON a.city_iata = c.iata;
