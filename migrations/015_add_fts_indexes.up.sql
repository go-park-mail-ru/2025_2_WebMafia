CREATE INDEX artist_fts_vector_idx ON artist USING GIN(fts_vector);
CREATE INDEX album_fts_vector_idx ON album USING GIN(fts_vector);
CREATE INDEX track_fts_vector_idx ON track USING GIN(fts_vector);