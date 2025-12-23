DROP INDEX IF EXISTS idx_playlist_title;
DROP INDEX IF EXISTS idx_playlist_is_favorite;
DROP INDEX IF EXISTS idx_playlist_created_at;

DROP TABLE IF EXISTS playlist;

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

DROP TRIGGER IF EXISTS update_playlist_updated_at ON playlist;

CREATE TRIGGER update_playlist_updated_at
    BEFORE UPDATE ON playlist
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

DROP INDEX IF EXISTS idx_playlist_track_track_id;
DROP INDEX IF EXISTS idx_playlist_track_playlist_id;

DROP TABLE IF EXISTS playlist_track;

CREATE TABLE playlist_track (
    playlist_id UUID NOT NULL,
    track_id UUID NOT NULL,

    PRIMARY KEY (playlist_id, track_id),
    CONSTRAINT fk_playlist_track_playlist
        FOREIGN KEY (playlist_id)
            REFERENCES playlist(playlist_id)
            ON DELETE CASCADE,

    CONSTRAINT fk_playlist_track_track
        FOREIGN KEY (track_id)
            REFERENCES track(track_id)
            ON DELETE CASCADE
);

COMMENT ON TABLE playlist_track IS 'Связь треков и плейлистов';