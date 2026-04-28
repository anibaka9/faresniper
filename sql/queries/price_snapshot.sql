-- name: CreatePriceSnapshot :exec
INSERT
    OR IGNORE INTO price_snapshots(
        flight_number,
        origin_iata,
        destination_iata,
        departure_at,
        airline_iata,
        price,
        observed_at
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?);
