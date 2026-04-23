-- +goose Up
CREATE TABLE cities (
    iata TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    country_code TEXT NOT NULL,
    FOREIGN KEY (country_code) REFERENCES countries(code) ON DELETE CASCADE,
);

-- +goose Down
DROP TABLE cities;
