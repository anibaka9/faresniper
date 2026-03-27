-- +goose Up
CREATE TABLE cities (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    country_id ITEGER NOT NULL,
    FOREIGN KEY (country_id) REFERENCES countries(id) ON DELETE CASCADE,
    UNIQUE (name, country_id)
);

-- +goose Down
DROP TABLE cities;
