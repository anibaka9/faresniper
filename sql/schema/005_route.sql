-- +goose Up
CREATE TABLE routes (
    id INTEGER PRIMARY KEY,
    flightsfrom_id INTEGER UNIQUE NOT NULL,
    airline_id INTEGER NOT NULL,
    airport_from_id INTEGER NOT NULL,
    airport_to_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL,
    FOREIGN KEY (airline_id) REFERENCES airlines(id) ON DELETE CASCADE,
    FOREIGN KEY (airport_from_id) REFERENCES airports(id) ON DELETE CASCADE,
    FOREIGN KEY (airport_to_id) REFERENCES airports(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE routes;
