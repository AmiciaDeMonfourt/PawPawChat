CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL,
    hash_pass TEXT NOT NULL
);

CREATE TABLE user_info (
    id BIGINT PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    phone VARCHAR(11) DEFAULT '-',
    country VARCHAR(32) DEFAULT '-',
    city VARCHAR(32) DEFAULT '-',
    age INTEGER DEFAULT -1,
    birthday DATE DEFAULT NULL
);

CREATE TABLE user_avatars (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    display_order INTEGER NOT NULL,
    url VARCHAR(256) NOT NULL
);

CREATE TABLE user_friends (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    friend_id BIGINT NOT NULL REFERENCES users(id),
    UNIQUE(user_id, friend_id)
);

CREATE TABLE user_subscribers (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    subscriber_id BIGINT NOT NULL REFERENCES users(id),
    UNIQUE(user_id, subscriber_id)
);
