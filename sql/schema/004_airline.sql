-- +goose Up
CREATE TABLE airlines (
    id INTEGER PRIMARY KEY,
    callsign TEXT NOT NULL,
    icao TEXT NOT NULL,
    iata TEXT NOT NULL,
    name TEXT NOT NULL,
    short_name TEXT NOT NULL,
    full_name TEXT NOT NULL,
    is_active BOOLEAN NOT NULL,
    country_id INTEGER,
    FOREIGN KEY (country_id) REFERENCES countries(id)
);

-- +goose Down
DROP TABLE airlines;
