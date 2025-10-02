CREATE TABLE track_artist (
    track_id UUID NOT NULL,
    artist_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (track_id, artist_id),
    CONSTRAINT fk_track_artist_track
        FOREIGN KEY (track_id)
        REFERENCES track(track_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_track_artist_artist
        FOREIGN KEY (artist_id)
        REFERENCES artist(artist_id)
        ON DELETE CASCADE
);

CREATE TABLE track_genre (
    track_id UUID NOT NULL,
    genre_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (track_id, genre_id),
    CONSTRAINT fk_track_genre_track
        FOREIGN KEY (track_id)
        REFERENCES track(track_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_track_genre_genre
        FOREIGN KEY (genre_id)
        REFERENCES genre(genre_id)
        ON DELETE CASCADE
);

CREATE TABLE track_album (
    track_id UUID NOT NULL,
    album_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (track_id, album_id),
    CONSTRAINT fk_track_album_track
        FOREIGN KEY (track_id)
        REFERENCES track(track_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_track_album_album
        FOREIGN KEY (album_id)
        REFERENCES album(album_id)
        ON DELETE CASCADE
);

CREATE TABLE track_playlist (
    track_id UUID NOT NULL,
    playlist_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (track_id, playlist_id),
    CONSTRAINT fk_track_playlist_track
        FOREIGN KEY (track_id)
        REFERENCES track(track_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_track_playlist_playlist
        FOREIGN KEY (playlist_id)
        REFERENCES playlist(playlist_id)
        ON DELETE CASCADE
);

CREATE INDEX idx_track_artist_artist_id ON track_artist (artist_id);
CREATE INDEX idx_track_genre_genre_id ON track_genre (genre_id);
CREATE INDEX idx_track_album_album_id ON track_album (album_id);
CREATE INDEX idx_track_playlist_playlist_id ON track_playlist (playlist_id);