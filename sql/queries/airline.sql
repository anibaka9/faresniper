-- name: CreateAirline :one
INSERT INTO
    airlines (iata, name, is_lowcost)
VALUES
    (?, ?, ?)
RETURNING
    *;

-- name: GetAirline :one
SELECT
    *
FROM
    airlines
WHERE
    iata = ?
LIMIT
    1;

-- name: GetAirlines :many
SELECT
    *
FROM
    airlines
ORDER BY
    name
LIMIT
    ? OFFSET ?;

-- name: CountAirlines :one
SELECT
    COUNT(*) AS count_airlines
FROM
    airlines;

-- name: UpdateAirline :one
UPDATE
    airlines
SET
    name = ?,
    is_lowcost = ?
WHERE
    iata = ?
RETURNING
    *;

-- name: GetAllAirlines :many
SELECT
    iata,
    name
FROM
    airlines;
