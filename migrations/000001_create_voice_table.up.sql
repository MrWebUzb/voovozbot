CREATE TABLE IF NOT EXISTS voices (
    file_unique_id VARCHAR PRIMARY KEY,
    file_id VARCHAR,
    duration BIGINT,
    mime_type VARCHAR,
    file_size BIGINT,
    caption VARCHAR,
    created_at timestamp default now()
);

CREATE INDEX idx_file_id ON voices (file_id);
CREATE INDEX idx_caption ON voices (caption);