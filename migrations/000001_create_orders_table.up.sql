CREATE TABLE IF NOT EXISTS order_models (
    id VARCHAR(255) PRIMARY KEY,
    price DECIMAL(10, 2) NOT NULL,
    tax DECIMAL(10, 2) NOT NULL,
    final_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL
);
