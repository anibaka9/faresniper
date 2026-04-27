-- +goose Up
CREATE TABLE airlines (
    iata TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    is_lowcost BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE airlines;
