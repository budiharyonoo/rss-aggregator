-- +goose Up

CREATE TABLE feeds
(
    id         UUID PRIMARY KEY,
    user_id    UUID REFERENCES users (id) ON DELETE CASCADE,
    name       TEXT        NOT NULL,
    URL        TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP   NULL,
    updated_at TIMESTAMP   NULL
);

-- +goose Down

DROP TABLE feeds;
