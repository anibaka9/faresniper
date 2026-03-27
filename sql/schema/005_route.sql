-- +goose Up
CREATE TABLE routes (
    id INTEGER PRIMARY KEY,
    airline_id INTEGER NOT NULL,
    airport_from_id INTEGER NOT NULL,
    airport_to_id INTEGER NOT NULL,
    is_active BOOLEAN NOT NULL,
    day_1 BOOLEAN NOT NULL,
    day_2 BOOLEAN NOT NULL,
    day_3 BOOLEAN NOT NULL,
    day_4 BOOLEAN NOT NULL,
    day_5 BOOLEAN NOT NULL,
    day_6 BOOLEAN NOT NULL,
    day_7 BOOLEAN NOT NULL,
    FOREIGN KEY (airline_id) REFERENCES airlines(id) ON DELETE CASCADE,
    FOREIGN KEY (airport_from_id) REFERENCES airports(id) ON DELETE CASCADE,
    FOREIGN KEY (airport_to_id) REFERENCES airports(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE routes;
