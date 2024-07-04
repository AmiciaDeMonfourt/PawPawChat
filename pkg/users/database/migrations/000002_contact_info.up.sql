CREATE TABLE contact_info (
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    phone   VARCHAR(11) DEFAULT '-',
    country VARCHAR(32) DEFAULT '-',
    city    VARCHAR(32) DEFAULT '-',
    PRIMARY KEY (user_id)
);
