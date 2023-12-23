-- +goose Up

CREATE TABLE users
(
    id         UUID PRIMARY KEY,
    name       varchar(255) NOT NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL
);

-- +goose Down

DROP TABLE users;
