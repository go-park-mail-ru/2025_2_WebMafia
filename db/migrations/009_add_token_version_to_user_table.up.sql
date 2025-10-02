ALTER TABLE "user"
ADD COLUMN token_version INTEGER NOT NULL DEFAULT 1,
ADD COLUMN last_logout_at TIMESTAMPTZ;

COMMENT ON COLUMN "user".token_version IS 'Версия токена для инвалидации JWT';
COMMENT ON COLUMN "user".last_logout_at IS 'Время последнего выхода из системы';

CREATE INDEX idx_user_token_version ON "user" (token_version);