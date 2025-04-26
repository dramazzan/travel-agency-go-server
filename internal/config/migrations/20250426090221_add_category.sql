-- +goose Up
ALTER TABLE tours ADD COLUMN category VARCHAR(255) DEFAULT 'new' NOT NULL;

-- +goose Down
ALTER TABLE tours DROP COLUMN category;
