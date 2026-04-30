CREATE TABLE tickets (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    reference_no VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL,
    user_id BIGINT
);