CREATE TABLE auth_info (
    user_id    BIGINT REFERENCES users(id) ON DELETE CASCADE,
    email      VARCHAR(64) NOT NULL UNIQUE,
    hash_pass  TEXT NOT NULL,
    last_login TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);

CREATE UNIQUE INDEX idx_auth_info_email ON auth_info(email);
