-- +goose Up
CREATE TABLE authors (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    bio TEXT
);

-- +goose Down
DROP TABLE authors;
