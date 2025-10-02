CREATE TABLE artist (
    artist_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    artist_name TEXT NOT NULL,
    avatar_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE artist IS 'Таблица хранения данных исполнителя';

CREATE INDEX idx_artist_name ON artist (artist_name);
CREATE INDEX idx_artist_created_at ON artist (created_at);