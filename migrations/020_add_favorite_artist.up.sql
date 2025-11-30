CREATE TABLE favorite_artist (
                                 user_id UUID NOT NULL,
                                 artist_id UUID NOT NULL,
                                 created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

                                 PRIMARY KEY (user_id, artist_id),

                                 CONSTRAINT fk_fav_artist_user
                                     FOREIGN KEY (user_id)
                                         REFERENCES "user"(user_id)
                                         ON DELETE CASCADE,

                                 CONSTRAINT fk_fav_artist_artist
                                     FOREIGN KEY (artist_id)
                                         REFERENCES artist(artist_id)
                                         ON DELETE CASCADE
);

COMMENT ON TABLE favorite_artist IS 'Избранные артисты пользователя';

CREATE INDEX idx_fav_artist_user ON favorite_artist (user_id);
CREATE INDEX idx_fav_artist_artist ON favorite_artist (artist_id);