CREATE OR REPLACE FUNCTION generate_fts_vector()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_TABLE_NAME = 'artist' THEN
        NEW.fts_vector := to_tsvector('simple', coalesce(NEW.artist_name, '') || ' ' || coalesce(NEW.description, ''));
    ELSIF TG_TABLE_NAME = 'album' THEN
        NEW.fts_vector := to_tsvector('simple', coalesce(NEW.title, '') || ' ' || coalesce(NEW.description, ''));
    ELSIF TG_TABLE_NAME = 'track' THEN
        NEW.fts_vector := to_tsvector('simple', coalesce(NEW.title, '') || ' ' || coalesce(NEW.description, ''));
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER artist_fts_update_trigger
BEFORE INSERT OR UPDATE ON artist
FOR EACH ROW EXECUTE FUNCTION generate_fts_vector();

CREATE TRIGGER album_fts_update_trigger
BEFORE INSERT OR UPDATE ON album
FOR EACH ROW EXECUTE FUNCTION generate_fts_vector();

CREATE TRIGGER track_fts_update_trigger
BEFORE INSERT OR UPDATE ON track
FOR EACH ROW EXECUTE FUNCTION generate_fts_vector();