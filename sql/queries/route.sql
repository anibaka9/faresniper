-- name: CreateRoute :one
INSERT INTO
    routes (
        flightsfrom_id,
        airline_id,
        airport_from_id,
        airport_to_id,
        is_active
    )
VALUES
    (?, ?, ?, ?, ?)
RETURNING
    *;

-- name: GetRouteByFlightsFromId :one
SELECT
    *
FROM
    routes
WHERE
    routes.flightsfrom_id = ?
LIMIT
    1;

-- name: GetRoute :one
SELECT
    r.id,
    r.flightsfrom_id,
    r.is_active,
    r.airline_id,
    ar.name AS airline_name,
    ar.iata AS airline_iata,
    r.airport_from_id,
    af.iata AS airport_from_iata,
    cf.name AS city_from_name,
    r.airport_to_id,
    at.iata AS airport_to_iata,
    ct.name AS city_to_name
FROM
    routes r
    JOIN airlines ar ON ar.id = r.airline_id
    JOIN airports af ON af.id = r.airport_from_id
    JOIN airports at ON at.id = r.airport_to_id
    JOIN cities cf ON cf.id = af.city_id
    JOIN cities ct ON ct.id = at.city_id
WHERE
    r.id = ?
LIMIT
    1;

-- name: GetRoutes :many
SELECT
    r.id,
    r.airport_from_id,
    cf.name AS city_from_name,
    af.iata AS airport_from_iata,
    r.airport_to_id,
    ct.name AS city_to_name,
    at.iata AS airport_to_iata,
    r.airline_id,
    ar.name AS airline_name
FROM
    routes r
    JOIN airports af ON af.id = r.airport_from_id
    JOIN airports at ON at.id = r.airport_to_id
    JOIN airlines ar ON ar.id = r.airline_id
    JOIN cities cf ON af.city_id = cf.id
    JOIN cities ct ON at.city_id = ct.id
LIMIT
    ? OFFSET ?;

-- name: CountRoutes :one
SELECT
    COUNT(*) AS count_routes
FROM
    routes;

-- name: UpdateRoute :one
UPDATE
    routes
SET
    flightsfrom_id = ?,
    airline_id = ?,
    airport_from_id = ?,
    airport_to_id = ?,
    is_active = ?
WHERE
    id = ?
RETURNING
    *;
