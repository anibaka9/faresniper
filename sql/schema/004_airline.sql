-- +goose Up
CREATE TABLE airlines (
    id INTEGER PRIMARY KEY,
    flightsfrom_id INTEGER UNIQUE NOT NULL,
    iata TEXT NOT NULL,
    name TEXT NOT NULL,
    is_active BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE airlines;
