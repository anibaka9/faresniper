-- name: CreateCountry :one
INSERT INTO
    countries (code, name)
VALUES
    (?, ?)
RETURNING
    *;

-- name: GetCountry :one
SELECT
    *
FROM
    countries
WHERE
    code = ?
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
    name = ?
WHERE
    code = ?
RETURNING
    *;

-- name: GetAllCountries :many
SELECT
    name,
    code
FROM
    countries;
