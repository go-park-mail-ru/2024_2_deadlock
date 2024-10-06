-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS article (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title TEXT NOT NULL,
    media_url TEXT,
    body TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS article;
-- +goose StatementEnd
