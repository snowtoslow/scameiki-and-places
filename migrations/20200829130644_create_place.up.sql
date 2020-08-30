CREATE TABLE IF NOT EXISTS places (
    id bigserial NOT NULL PRIMARY KEY,
    geolocation VARCHAR(255) NOT NULL,
    photo TEXT
);