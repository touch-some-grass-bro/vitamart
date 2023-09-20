CREATE TABLE IF NOT EXISTS items (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    image_binary BYTEA NOT NULL,
    price INT NOT NULL,
    seller_email VARCHAR(300) UNIQUE NOT NULL REFERENCES users(email) ON DELETE CASCADE,
    isSold BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);