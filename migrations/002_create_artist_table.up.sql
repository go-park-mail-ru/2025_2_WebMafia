CREATE TABLE artist (
    artist_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    artist_name TEXT NOT NULL,
    avatar_url TEXT,
    header_url TEXT,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_artist_name_length CHECK (length(artist_name) >= 1),
    CONSTRAINT chk_artist_avatar_url_extension CHECK (avatar_url IS NULL OR avatar_url LIKE '%.%')
);

COMMENT ON TABLE artist IS 'Таблица хранения данных исполнителя';
COMMENT ON COLUMN artist.description IS 'Описание исполнителя';

CREATE INDEX idx_artist_name ON artist (artist_name);
