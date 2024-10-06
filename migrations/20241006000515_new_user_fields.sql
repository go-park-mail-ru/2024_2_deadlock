-- +goose Up
-- +goose StatementBegin
ALTER TABLE account ADD COLUMN num_subscribers INTEGER NOT NULL DEFAULT 0,
ADD COLUMN num_subscriptions INTEGER NOT NULL DEFAULT 0,
ADD COLUMN registration_date DATE NOT NULL DEFAULT CURRENT_DATE,
ADD COLUMN extra_info TEXT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE account
DROP COLUMN num_subscribers,
DROP COLUMN num_subscriptions,
DROP COLUMN registration_date,
DROP COLUMN extra_info;
-- +goose StatementEnd
