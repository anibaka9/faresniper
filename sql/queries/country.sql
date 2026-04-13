-- name: CreateCountry :one
INSERT INTO
    countries (country_code, name)
VALUES
    (?, ?)
RETURNING
    *;

-- name: GetCountryByCode :one
SELECT
    *
FROM
    countries
WHERE
    country_code = ?
LIMIT
    1;

-- name: GetCountry :one
SELECT
    *
FROM
    countries
WHERE
    id = ?
LIMIT
    1;

-- name: GetCountries :many
SELECT
    *
FROM
    countries
LIMIT
    ? OFFSET ?;

-- name: CountCountries :one
SELECT
    COUNT(*) AS count_countries
FROM
    countries;

-- name: UpdateCountry :one
UPDATE
    countries
SET
    country_code = ?,
    name = ?
WHERE
    id = ?
RETURNING
    *;
