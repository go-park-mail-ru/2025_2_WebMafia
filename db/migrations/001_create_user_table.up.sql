CREATE TABLE "user" (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    avatar_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE "user" IS 'Таблица хранения учетных записей пользователей';

CREATE INDEX idx_user_login ON "user" (login);
CREATE INDEX idx_user_email ON "user" (email);
CREATE INDEX idx_user_created_at ON "user" (created_at);