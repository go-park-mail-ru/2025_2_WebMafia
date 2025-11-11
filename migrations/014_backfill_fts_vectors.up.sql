UPDATE artist
SET fts_vector = to_tsvector('simple', coalesce(artist_name, '') || ' ' || coalesce(description, ''));

UPDATE album
SET fts_vector = to_tsvector('simple', coalesce(title, '') || ' ' || coalesce(description, ''));

UPDATE track
SET fts_vector = to_tsvector('simple', coalesce(title, '') || ' ' || coalesce(description, ''));