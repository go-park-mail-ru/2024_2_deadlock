-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS feed;

CREATE TABLE IF NOT EXISTS feed.article (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title TEXT,
    media_url TEXT,
    body TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS feed.article;
DROP SCHEMA IF EXISTS feed;
-- +goose StatementEnd
