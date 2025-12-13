CREATE TABLE comment (
    comment_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    track_id UUID NOT NULL,
    user_id UUID NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_comment_track
        FOREIGN KEY (track_id)
        REFERENCES track(track_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_comment_user
        FOREIGN KEY (user_id)
        REFERENCES "user"(user_id)
        ON DELETE CASCADE,

    CONSTRAINT chk_comment_text_length CHECK (length(text) > 0 AND length(text) <= 1000)
);

CREATE INDEX idx_comment_track_id_created_at ON comment (track_id, created_at DESC);
