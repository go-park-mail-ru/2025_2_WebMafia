CREATE TABLE "user" (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    avatar_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
    CONSTRAINT chk_user_login_length CHECK (length(login) <= 30),
    CONSTRAINT chk_user_email_format CHECK (email LIKE '%@%'),
    CONSTRAINT chk_user_avatar_url_extension CHECK (avatar_url = '' OR avatar_url LIKE '%.%')
);

COMMENT ON TABLE "user" IS 'Таблица хранения учетных записей пользователей';