-- +goose Up
-- +goose StatementBegin
ALTER TABLE feed.article ADD COLUMN author_id INTEGER;
UPDATE feed.article SET author_id = 1;
ALTER TABLE feed.article ALTER COLUMN author_id SET NOT NULL;
ALTER TABLE feed.article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES auth.user(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feed.article DROP COLUMN author_id;
-- +goose StatementEnd
