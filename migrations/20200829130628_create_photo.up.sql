CREATE TABLE IF NOT EXISTS photo (
    id bigserial NOT NULL PRIMARY KEY,
    photo TEXT NOT NULL,
    createdAt Date,
    updatedAt Date
);