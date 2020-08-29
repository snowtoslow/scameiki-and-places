
CREATE TABLE IF NOT EXISTS places (
    id bigserial NOT NULL PRIMARY KEY,
    geolocation VARCHAR(255) NOT NULL,
    createdAt Date,
    updatedAt Date,
    photoId bigserial NOT NULL
);