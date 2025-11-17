CREATE TABLE playlist (
    playlist_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    avatar_url TEXT,
    is_favorite BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_playlist_user
        FOREIGN KEY (user_id)
        REFERENCES "user"(user_id)
        ON DELETE CASCADE,
    CONSTRAINT chk_playlist_title_length CHECK (length(title) >= 1)
);

COMMENT ON TABLE playlist IS 'Плейлисты пользователей';

CREATE INDEX idx_playlist_user_id ON playlist (user_id);
CREATE INDEX idx_playlist_created_at ON playlist (created_at);
