CREATE TRIGGER update_playlist_updated_at
    BEFORE UPDATE ON playlist
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
