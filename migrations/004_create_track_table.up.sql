CREATE TABLE track (
    track_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    duration_s INTEGER NOT NULL CHECK (duration_s > 0),
    file_url TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_title_length CHECK (length(title) >= 1)
);

COMMENT ON TABLE track IS 'Таблица хранения треков';

CREATE INDEX idx_track_title ON track (title);
