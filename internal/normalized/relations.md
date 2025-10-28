# Схема данных веб-сервиса WaveMusic

## Отношения и зависимости:

### Пользователи и профили:

#### Relation: 'user'
**Описание:** Таблица хранения учетных записей пользователей <br>
**Отношение:** <br>
USER (

    USER_ID UUID PRIMARY KEY, 
    LOGIN TEXT NOT NULL UNIQUE, 
    PASSWORD_HASH TEXT NOT NULL, 
    EMAIL TEXT NOT NULL UNIQUE, 
    AVATAR_URL TEXT,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP 
)<br>
**Функциональные зависимости:** <br>
{USER_ID} → LOGIN, PASSWORD_HASH, EMAIL, CREATED_AT, UPDATED_AT <br>
{LOGIN} → USER_ID, PASSWORD_HASH, EMAIL, CREATED_AT, UPDATED_AT <br>
{EMAIL} → USER_ID, LOGIN, PASSWORD_HASH, CREATED_AT, UPDATED_AT <br>

### Каталог:

#### Relation: 'artist'
**Описание:** Таблица хранения данных исполнителя <br>
**Отношение:** <br>
ARTIST (

    ARTIST_ID UUID PRIMARY KEY, 
    ARTIST_NAME TEXT NOT NULL, 
    AVATAR_URL TEXT, 
    DESCRIPTION TEXT,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
)<br>
**Функциональные зависимости:** <br>
{ARTIST_ID} → ARTIST_NAME, AVATAR_URL, CREATED_AT, UPDATED_AT <br>

#### Relation: 'genre'
**Описание:** Таблица хранения жанров <br>
**Отношение:** <br>
GENRE (

    GENRE_ID UUID PRIMARY KEY, 
    GENRE_NAME TEXT NOT NULL UNIQUE, 
    DESCRIPTION TEXT,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP 
) <br>
**Функциональные зависимости:** <br>
{GENRE_ID} → GENRE_NAME, CREATED_AT <br>

#### Relation: 'track'
**Описание:** Таблица хранения треков <br>
**Отношение:** <br>
TRACK (

    TRACK_ID UUID PRIMARY KEY, 
    TITLE TEXT NOT NULL, 
    DURATION_MS INTEGER NOT NULL,
    FILE_URL TEXT NOT NULL, 
    DESCRIPTION TEXT,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
) <br>
**Функциональные зависимости:**<br>
{TRACK_ID} → TITLE, DURATION_MS, FILE_URL, CREATED_AT, UPDATED_AT <br>

#### Relation: 'album'
**Описание:** Таблица хранения альбомов <br>
**Отношение:** <br>
ALBUM (

    ALBUM_ID UUID PRIMARY KEY, 
    TITLE TEXT NOT NULL, 
    AVATAR_URL TEXT, 
    ARTIST_ID UUID NOT NULL, 
    DESCRIPTION TEXT,
    RELEASE_DATE DATE, 
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP 
) <br>
**Функциональные зависимости:** <br>
{ALBUM_ID} → TITLE, AVATAR_URL, ARTIST_ID, RELEASE_DATE, CREATED_AT, UPDATED_AT <br>

#### Relation: 'playlist'
**Описание:** Таблица хранения плейлистов <br>
**Отношение:** <br>
PLAYLIST (

    PLAYLIST_ID UUID PRIMARY KEY, 
    TITLE TEXT NOT NULL, 
    AVATAR_URL TEXT, 
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    UPDATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    USER_ID UUID NOT NULL 
) <br>
**Функциональные зависимости:** <br>
{PLAYLIST_ID} → TITLE, AVATAR_URL, CREATED_AT, UPDATED_AT, USER_ID <br>

### Связующие таблицы:

#### Relation: 'track to artist'
**Описание:** Таблица реализации many-to-many связи между треками и исполнителями <br>
**Отношение:**<br>
TRACK_ARTIST (

    TRACK_ID UUID NOT NULL, 
    ARTIST_ID UUID NOT NULL, 
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    PRIMARY KEY (TRACK_ID, ARTIST_ID) 
) <br>
**Функциональные зависимости:** <br>
{TRACK_ID, ARTIST_ID} → CREATED_AT <br>

#### Relation: 'track to genre'
**Описание:** Таблица реализации many-to-many связи между треками и жанрами <br>
**Отношение:** <br>
TRACK_GENRE (

    TRACK_ID UUID NOT NULL,
    GENRE_ID UUID NOT NULL,
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (TRACK_ID, GENRE_ID)
) <br>
**Функциональные зависимости:** <br>
{TRACK_ID, GENRE_ID} → CREATED_AT <br>

#### Relation: 'album to track'
**Описание:** Таблица реализации many-to-many связи между треками и альбомами <br>
**Отношение:** <br>
TRACK_ALBUM (

    TRACK_ID UUID NOT NULL, 
    ALBUM_ID UUID NOT NULL, 
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    PRIMARY KEY (TRACK_ID, ALBUM_ID) 
) <br>
**Функциональные зависимости:** <br>
{TRACK_ID, ALBUM_ID} → CREATED_AT <br>

#### Relation: 'playlist to track'
**Описание:** Таблица реализации many-to-many связи между треками и плейлистами <br>
**Отношение:** <br>
TRACK_PLAYLIST (

    TRACK_ID UUID NOT NULL, 
    PLAYLIST_ID UUID NOT NULL, 
    CREATED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    PRIMARY KEY (TRACK_ID, PLAYLIST_ID) 
) <br>
**Функциональные зависимости:** <br>
{TRACK_ID, PLAYLIST_ID} → CREATED_AT <br>

#### Relation: 'user liked track'
**Описание:** Хранит треки, понравившиеся пользователю <br>
**Отношение:** <br>
USER_LIKED_TRACK (

    USER_ID UUID NOT NULL, 
    TRACK_ID UUID NOT NULL, 
    LIKED_AT TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, 
    PRIMARY KEY (USER_ID, TRACK_ID), 
) <br>
**Функциональные зависимости:** <br>
{USER_ID, TRACK_ID} → LIKED_AT <br>


## Соответствие требованиям:

### Соответствие 1НФ:
Схема соответствует 1 нормальной форме, так как: <br>
- Каждый из атрибутов атомарен <br>
- Каждое значение в одном столбце имеет один и тот же тип <br>
- Порядок строк и столбцов не имеет значения <br>

### Соответствие 2НФ:
Схема соответствует 2 нормальной форме, так как: <br>
- Она соответствует 1 нормальной форме <br>
- Неключевые атрибуты функционально зависят от всего первичного ключа (в таблицах `user`, `artist`, `genre`, `track`, `playlist`, `album`) <br>
- В таблицах с составным ключом нет неключевых атрибутов, зависящих от части ключа (в таблицах `track_artist`, `track_genre`, `album_track`, `playlist_track`) <br>
- В таблицах с составным ключом и дополнительными атрибутами зависят от всего составного ключа (в таблице `user_liked_tracks` атрибут `liked_at` определен только в контексте конкретного трека) <br>

### Соответствие 3НФ:
Схема соответствует 3 нормальной форме, так как: <br>
- Соответствует 1 и 2 нормальным формам <br>
- Нет транзитивных зависимостей между неключевыми атрибутами (в таблице `album` атрибут `avatar_url` зависит только от `album_id`, в остальном аналогично) <br>

### Соответствие НФБК:
Схема соответствует нормальной форме Бойса-Кодда, так как: <br>
- Соответствует 1, 2 и 3 нормальным формам <br>
- Содержит потенциальные ключи (в таблицах: `user_id`, `profile_id`, `artist_id`, в остальном аналогично)<br>
  
