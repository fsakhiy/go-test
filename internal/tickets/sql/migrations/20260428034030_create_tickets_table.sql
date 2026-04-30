-- +goose Up
-- +goose StatementBegin
CREATE TABLE tickets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    reference_no VARCHAR(100) NOT NULL UNIQUE,
    status VARCHAR(50) NOT NULL DEFAULT 'open',
    user_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tickets;
-- +goose StatementEnd