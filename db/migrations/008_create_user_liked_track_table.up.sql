CREATE TABLE user_liked_track (
    user_id UUID NOT NULL,
    track_id UUID NOT NULL,
    liked_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, track_id),
    CONSTRAINT fk_user_liked_track_user
        FOREIGN KEY (user_id)
        REFERENCES "user"(user_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_user_liked_track_track
        FOREIGN KEY (track_id)
        REFERENCES track(track_id)
        ON DELETE CASCADE
);

COMMENT ON TABLE user_liked_track IS 'Треки, понравившиеся пользователю';

CREATE INDEX idx_user_liked_track_user_id ON user_liked_track (user_id);
CREATE INDEX idx_user_liked_track_track_id ON user_liked_track (track_id);
CREATE INDEX idx_user_liked_track_liked_at ON user_liked_track (liked_at);