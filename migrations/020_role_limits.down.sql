ALTER ROLE auth_user RESET statement_timeout;
ALTER ROLE auth_user RESET lock_timeout;
ALTER ROLE auth_user RESET idle_in_transaction_session_timeout;
ALTER ROLE auth_user RESET temp_file_limit;
ALTER ROLE auth_user RESET search_path;

ALTER ROLE catalog_user RESET statement_timeout;
ALTER ROLE catalog_user RESET lock_timeout;
ALTER ROLE catalog_user RESET idle_in_transaction_session_timeout;
ALTER ROLE catalog_user RESET temp_file_limit;
ALTER ROLE catalog_user RESET search_path;

ALTER ROLE playlist_user RESET statement_timeout;
ALTER ROLE playlist_user RESET lock_timeout;
ALTER ROLE playlist_user RESET idle_in_transaction_session_timeout;
ALTER ROLE playlist_user RESET temp_file_limit;
ALTER ROLE playlist_user RESET search_path;
