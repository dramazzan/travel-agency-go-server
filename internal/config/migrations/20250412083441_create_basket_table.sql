-- +goose Up
CREATE TABLE baskets (
                         id SERIAL PRIMARY KEY,
                         user_id INT UNIQUE NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         deleted_at TIMESTAMP,
                         CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE basket_tours (
                              basket_id INT NOT NULL,
                              tour_id INT NOT NULL,
                              PRIMARY KEY (basket_id, tour_id),
                              FOREIGN KEY (basket_id) REFERENCES baskets(id) ON DELETE CASCADE,
                              FOREIGN KEY (tour_id) REFERENCES tours(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS basket_tours;
DROP TABLE IF EXISTS baskets;
