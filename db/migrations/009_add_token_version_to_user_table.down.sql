ALTER TABLE "user"
DROP COLUMN IF EXISTS token_version,
DROP COLUMN IF EXISTS last_logout_at;