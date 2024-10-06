-- +goose Up
-- +goose StatementBegin
ALTER TABLE article ADD COLUMN author_id INTEGER NOT NULL DEFAULT 1;
ALTER TABLE article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES account(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE article DROP COLUMN author_id;
-- +goose StatementEnd
