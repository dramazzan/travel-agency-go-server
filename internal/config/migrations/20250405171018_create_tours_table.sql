-- +goose Up
CREATE TABLE tours (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       description TEXT,
                       destination VARCHAR(255) NOT NULL,
                       start_date TIMESTAMP NOT NULL,
                       end_date TIMESTAMP NOT NULL,
                       price FLOAT CHECK(price >= 0) NOT NULL,
                       max_capacity INT,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS tours;
