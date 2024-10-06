-- +goose Up
-- +goose StatementBegin
ALTER TABLE account 
ADD COLUMN avatar_url TEXT NULL,
ADD COLUMN first_name TEXT NOT NULL DEFAULT '',
ADD COLUMN last_name TEXT NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE account
DROP COLUMN avatar_url,
DROP COLUMN first_name,
DROP COLUMN last_name;
-- +goose StatementEnd
