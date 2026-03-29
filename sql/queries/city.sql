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
