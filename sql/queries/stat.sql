-- name: GetStats :one
SELECT
    (
        SELECT
            COUNT(*)
        FROM
            airlines
    ) AS airlines_num,
    (
        SELECT
            COUNT(*)
        FROM
            airports
    ) AS airports_num,
    (
        SELECT
            COUNT(*)
        FROM
            cities
    ) AS cities_num,
    (
        SELECT
            COUNT(*)
        FROM
            countries
    ) AS countries_num;
