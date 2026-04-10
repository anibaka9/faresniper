-- name: CreateCity :one
INSERT INTO
    cities (name, country_id)
VALUES
    (?, ?)
RETURNING
    *;

-- name: GetCityByCountyCodeAndName :one
SELECT
    cities.*
FROM
    cities
    JOIN countries ON cities.country_id = countries.id
WHERE
    cities.name = ?
    AND countries.country_code = ?
LIMIT
    1;

-- name: GetCities :many
SELECT
    ct.*,
    cr.name country_name
FROM
    cities ct
    JOIN countries cr ON ct.country_id = cr.id
LIMIT
    ? OFFSET ?;

-- name: CountCities :one
SELECT
    COUNT(*) AS count_cities
FROM
    cities;
