CREATE TABLE media_metadata (
    id           BIGSERIAL PRIMARY KEY,
    media_id     BIGINT REFERENCES users_avatars(id) ON DELETE CASCADE,
    file_name    VARCHAR(255) NOT NULL,
    file_ext     VARCHAR(10) NOT NULL,
    file_size    BIGINT,
    content_type VARCHAR(30),
    duration     INTEGER,
    resolution   VARCHAR(20),
    uploaded_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_media_metadata_media_id ON media_metadata(media_id);
