-- +goose Up
CREATE TABLE airports (
    id INTEGER PRIMARY KEY,
    iata TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    city_id INTEGER NOT NULL,
    FOREIGN KEY (city_id) REFERENCES cities(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE airports;
