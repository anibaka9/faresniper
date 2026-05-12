-- +goose Up
CREATE TABLE watched_routes (
    origin_iata TEXT NOT NULL,
    destination_iata TEXT NOT NULL,
    active BOOLEAN NOT NULL,
    PRIMARY KEY (origin_iata, destination_iata),
    FOREIGN KEY (origin_iata) REFERENCES airports(iata),
    FOREIGN KEY (destination_iata) REFERENCES airports(iata)
);

-- +goose Down
DROP TABLE watched_routes;
