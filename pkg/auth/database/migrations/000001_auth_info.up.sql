CREATE TABLE auth_info (
    user_id    BIGSERIAL PRIMARY KEY,
    email      VARCHAR(64) NOT NULL UNIQUE,
    hash_pass  TEXT NOT NULL,
    last_login TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);

CREATE UNIQUE INDEX idx_auth_info_email ON auth_info(email);
