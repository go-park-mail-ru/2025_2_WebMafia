CREATE TABLE genre (
    genre_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    genre_name TEXT NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_genre_name_length CHECK (length(genre_name) >= 3)
);

COMMENT ON TABLE genre IS 'Таблица хранения жанров';

