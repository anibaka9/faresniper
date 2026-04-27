-- +goose Up
CREATE TABLE airports (
    iata TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    iata_type TEXT NOT NULL,
    flightable BOOLEAN NOT NULL,
    city_iata TEXT NOT NULL,
    FOREIGN KEY (city_iata) REFERENCES cities(iata) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE airports;
