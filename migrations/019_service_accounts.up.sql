-- auth
CREATE USER auth_user WITH PASSWORD 'auth_password';

GRANT CONNECT ON DATABASE mydb TO auth_user;
GRANT USAGE ON SCHEMA public TO auth_user;

GRANT SELECT, INSERT, UPDATE ON TABLE "user" TO auth_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE ON TABLES TO auth_user;

REVOKE CREATE ON SCHEMA public FROM auth_user;
REVOKE ALL ON ALL SEQUENCES IN SCHEMA public FROM auth_user;

-- service
CREATE USER catalog_user WITH PASSWORD 'catalog_password';

GRANT CONNECT ON DATABASE mydb TO catalog_user;
GRANT USAGE ON SCHEMA public TO catalog_user;

GRANT SELECT ON TABLE
    album,
    artist,
    track,
    track_artist,
    track_album,
    genre,
    track_genre
    TO catalog_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT ON TABLES TO catalog_user;

REVOKE CREATE ON SCHEMA public FROM catalog_user;
REVOKE ALL ON ALL SEQUENCES IN SCHEMA public FROM catalog_user;

-- playlist
CREATE USER playlist_user WITH PASSWORD 'playlist_password';

GRANT CONNECT ON DATABASE mydb TO playlist_user;
GRANT USAGE ON SCHEMA public TO playlist_user;

GRANT SELECT, INSERT, UPDATE, DELETE ON TABLE
    playlist,
    playlist_track
    TO playlist_user;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO playlist_user;

REVOKE CREATE ON SCHEMA public FROM playlist_user;
REVOKE ALL ON ALL SEQUENCES IN SCHEMA public FROM playlist_user;

