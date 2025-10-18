DROP TRIGGER IF EXISTS update_user_updated_at ON "user";
DROP TRIGGER IF EXISTS update_artist_updated_at ON artist;
DROP TRIGGER IF EXISTS update_track_updated_at ON track;
DROP TRIGGER IF EXISTS update_album_updated_at ON album;

DROP FUNCTION IF EXISTS update_updated_at_column();