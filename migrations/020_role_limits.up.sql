--auth
ALTER ROLE auth_user SET statement_timeout = '3s';
ALTER ROLE auth_user SET lock_timeout = '1s';
ALTER ROLE auth_user SET idle_in_transaction_session_timeout = '5s';
ALTER ROLE auth_user SET temp_file_limit = '50MB';
ALTER ROLE auth_user SET search_path = public;

-- catalog
ALTER ROLE catalog_user SET statement_timeout = '2s';
ALTER ROLE catalog_user SET lock_timeout = '1s';
ALTER ROLE catalog_user SET idle_in_transaction_session_timeout = '5s';
ALTER ROLE catalog_user SET temp_file_limit = '100MB';
ALTER ROLE catalog_user SET search_path = public;

-- playlist
ALTER ROLE playlist_user SET statement_timeout = '4s';
ALTER ROLE playlist_user SET lock_timeout = '1s';
ALTER ROLE playlist_user SET idle_in_transaction_session_timeout = '5s';
ALTER ROLE playlist_user SET temp_file_limit = '100MB';
ALTER ROLE playlist_user SET search_path = public;
