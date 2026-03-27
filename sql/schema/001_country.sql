-- +goose Up
CREATE TABLE countries (
    id INTEGER PRIMARY KEY,
    country_code TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE countries;
