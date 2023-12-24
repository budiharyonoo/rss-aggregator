-- +goose Up
CREATE TABLE posts
(
    id           UUID PRIMARY KEY,
    feed_id      UUID         NOT NULL REFERENCES feeds (id) ON DELETE CASCADE,
    title        VARCHAR(255) NOT NULL,
    url          TEXT UNIQUE  NOT NULL,
    description  TEXT         NULL,
    published_at TIMESTAMP    NOT NULL,
    created_at   TIMESTAMP    NULL,
    updated_at   TIMESTAMP    NULL
);

-- +goose Down
DROP TABLE posts;
