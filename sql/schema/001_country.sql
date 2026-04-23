-- +goose Up
CREATE TABLE countries (
    country_code TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE countries;
