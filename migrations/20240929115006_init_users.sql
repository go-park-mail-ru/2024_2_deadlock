-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS auth;

CREATE TABLE IF NOT EXISTS auth.user (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    email TEXT
        UNIQUE
        CONSTRAINT email_length CHECK ( CHAR_LENGTH(email) <= 255 )
        NOT NULL,
    password TEXT
        CONSTRAINT password_length CHECK ( CHAR_LENGTH(password) <= 255 )
        NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth.user;
DROP SCHEMA IF EXISTS auth;
-- +goose StatementEnd
