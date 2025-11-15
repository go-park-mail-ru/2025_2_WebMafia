ALTER TABLE user
    ADD COLUMN role TEXT NOT NULL CHECK (role IN ('user', 'admin', 'superadmin'))