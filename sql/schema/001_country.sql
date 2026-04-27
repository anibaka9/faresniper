-- +goose Up
CREATE TABLE countries (
    code TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE countries;
