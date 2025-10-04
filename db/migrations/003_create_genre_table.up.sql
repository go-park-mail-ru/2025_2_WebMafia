CREATE TABLE genre (
    genre_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    genre_name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE genre IS 'Таблица хранения жанров';

