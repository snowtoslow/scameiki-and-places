CREATE TABLE IF NOT EXISTS benches (
    id bigserial NOT NULL PRIMARY KEY,
    geolocation VARCHAR(255) NOT NULL,
    createdAt DATE,
    updatedDate DATE,
    photoId bigserial NOT NULL
);