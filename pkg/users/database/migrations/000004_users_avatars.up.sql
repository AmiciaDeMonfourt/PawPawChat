CREATE TABLE users_avatars (
    id       BIGSERIAL PRIMARY KEY,
    user_id  BIGINT REFERENCES users(id) ON DELETE CASCADE,
    bucket   VARCHAR(32),
    key      VARCHAR(32),
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    likes    INTEGER DEFAULT 0
);

CREATE UNIQUE INDEX idx_users_avatars_user_id ON users_avatars(user_id);