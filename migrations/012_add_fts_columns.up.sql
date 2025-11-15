ALTER TABLE artist ADD COLUMN fts_vector tsvector;
COMMENT ON COLUMN artist.fts_vector IS 'Поисковый вектор для полнотекстового поиска по исполнителю';

ALTER TABLE album ADD COLUMN fts_vector tsvector;
COMMENT ON COLUMN album.fts_vector IS 'Поисковый вектор для полнотекстового поиска по альбому';

ALTER TABLE track ADD COLUMN fts_vector tsvector;
COMMENT ON COLUMN track.fts_vector IS 'Поисковый вектор для полнотекстового поиска по треку';