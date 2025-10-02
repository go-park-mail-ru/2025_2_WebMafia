CREATE TABLE album (
    album_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    avatar_url TEXT,
    artist_id UUID NOT NULL,
    release_date DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_album_artist
        FOREIGN KEY (artist_id)
        REFERENCES artist(artist_id)
        ON DELETE CASCADE
);

COMMENT ON TABLE album IS 'Таблица хранения альбомов';

CREATE INDEX idx_album_title ON album (title);
CREATE INDEX idx_album_artist_id ON album (artist_id);
CREATE INDEX idx_album_release_date ON album (release_date);