CREATE TABLE favorite_album (
    user_id UUID NOT NULL,
    album_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id, album_id),

    CONSTRAINT fk_fav_album_user
        FOREIGN KEY (user_id)
            REFERENCES "user"(user_id)
            ON DELETE CASCADE,

    CONSTRAINT fk_fav_album_album
        FOREIGN KEY (album_id)
            REFERENCES album(album_id)
            ON DELETE CASCADE
);

COMMENT ON TABLE favorite_album IS 'Избранные альбомы пользователя';

CREATE INDEX idx_fav_album_user ON favorite_album (user_id);
CREATE INDEX idx_fav_album_album ON favorite_album (album_id);
