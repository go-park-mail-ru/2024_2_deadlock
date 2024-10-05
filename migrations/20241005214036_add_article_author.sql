-- +goose Up
-- +goose StatementBegin
ALTER TABLE article ADD COLUMN author_id INTEGER;
UPDATE article SET author_id = 1;
ALTER TABLE article ALTER COLUMN author_id SET NOT NULL;
ALTER TABLE article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES account(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE article DROP COLUMN author_id;
-- +goose StatementEnd
