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