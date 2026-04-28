-- +goose Up
CREATE TABLE price_snapshots (
    id INTEGER PRIMARY KEY,
    flight_number TEXT NOT NULL,
    origin_iata TEXT NOT NULL,
    destination_iata TEXT NOT NULL,
    departure_at TEXT NOT NULL,
    airline_iata TEXT NOT NULL,
    price NUMBER NOT NULL,
    observed_at TEXT NOT NULL,
    UNIQUE (
        flight_number,
        origin_iata,
        destination_iata,
        airline_iata,
        departure_at,
        observed_at
    ),
    FOREIGN KEY (origin_iata) REFERENCES airports(iata),
    FOREIGN KEY (destination_iata) REFERENCES airports(iata),
    FOREIGN KEY (airline_iata) REFERENCES airlines(iata)
);

-- +goose Down
DROP TABLE price_snapshots;
