ALTER TABLE "user"
    ADD COLUMN role TEXT NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin', 'superadmin'));

UPDATE "user" SET role = 'user' WHERE role IS NULL;