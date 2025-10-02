CREATE TABLE playlist (
    playlist_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    avatar_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id UUID NOT NULL,
    CONSTRAINT fk_playlist_user
        FOREIGN KEY (user_id)
        REFERENCES "user"(user_id)
        ON DELETE CASCADE
);

COMMENT ON TABLE playlist IS 'Таблица хранения плейлистов';

CREATE INDEX idx_playlist_title ON playlist (title);
CREATE INDEX idx_playlist_user_id ON playlist (user_id);