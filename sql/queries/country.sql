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

-- name: ListUsers :many
SELECT
    *
FROM
    countries;
