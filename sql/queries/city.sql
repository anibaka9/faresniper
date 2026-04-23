-- name: CreateCity :one
INSERT INTO
    cities (iata, name, country_code)
VALUES
    (?, ?, ?)
RETURNING
    *;

-- name: GetCity :one
SELECT
    ct.iata,
    ct.name,
    ct.country_code,
    cr.name AS country_name
FROM
    cities ct
    JOIN countries cr ON ct.country_code = cr.code
WHERE
    ct.iata = ?
LIMIT
    1;

-- name: GetCities :many
SELECT
    ct.iata,
    ct.name,
    ct.country_code,
    cr.name AS country_name,
FROM
    cities ct
    JOIN countries cr ON ct.country_code = cr.code
LIMIT
    ? OFFSET ?;

-- name: CountCities :one
SELECT
    COUNT(*) AS count_cities
FROM
    cities;

-- name: GetAllCities :many
SELECT
    ct.iata,
    cr.name || ' - ' || ct.name name
FROM
    cities ct
    JOIN countries cr ON ct.country_code = cr.code;

-- name: UpdateCity :one
UPDATE
    cities
SET
    name = ?,
    country_code = ?
WHERE
    iata = ?
RETURNING
    *;
