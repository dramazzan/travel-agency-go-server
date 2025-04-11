-- +goose Up
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(255) NOT NULL UNIQUE,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password TEXT NOT NULL,
                       role VARCHAR(255) DEFAULT 'user' NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP,
                       CONSTRAINT chk_role CHECK (role IN ('user', 'admin'))
);

-- +goose Down
DROP TABLE IF EXISTS users;
