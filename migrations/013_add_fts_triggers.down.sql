DROP TRIGGER IF EXISTS track_fts_update_trigger ON track;
DROP TRIGGER IF EXISTS album_fts_update_trigger ON album;
DROP TRIGGER IF EXISTS artist_fts_update_trigger ON artist;

DROP FUNCTION IF EXISTS generate_fts_vector();