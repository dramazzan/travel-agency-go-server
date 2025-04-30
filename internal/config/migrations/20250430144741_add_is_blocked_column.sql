-- +goose Up
ALTER TABLE users ADD COLUMN is_blocked BOOLEAN DEFAULT FALSE;

-- +goose Down
ALTER TABLE users DROP COLUMN is_blocked;
