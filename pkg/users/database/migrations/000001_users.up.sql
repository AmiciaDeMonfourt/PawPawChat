CREATE TABLE users (
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR(32) NOT NULL UNIQUE,
    first_name  VARCHAR(32) NOT NULL,
    second_name VARCHAR(32) NOT NULL,
    online      BOOLEAN DEFAULT TRUE,
    last_seen   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    age         INTEGER DEFAULT 0 CHECK (age >= 0),
    birthday    DATE DEFAULT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_blocked  BOOLEAN
);

CREATE UNIQUE INDEX idx_users_username ON users(username);
