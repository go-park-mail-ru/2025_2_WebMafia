import os
import json
from ytmusicapi import YTMusic
from datetime import datetime
import random
import spotipy
from spotipy.oauth2 import SpotifyClientCredentials
from minio import Minio
from upload_avatars import upload_avatar

minio_client = Minio(
    "localhost:8099",
    access_key="miniouser",
    secret_key="miniopassword",
    secure=False
)
bucket_name = "avatars"
if not minio_client.bucket_exists(bucket_name):
    minio_client.make_bucket(bucket_name)

client_id = "e2f9643cd6c24e50964226ebd093ac88"
client_secret = "723c43d6f74d445e8e5e592eed9a5833"

base_dir = os.path.dirname(os.path.abspath(__file__))
output_file = os.path.join(base_dir, "..", "migrations", "010_insert_API_data.up.sql")

# --- Подключение к YTMusic ---
yt = YTMusic(language='ru')

# --- Подключение к спотику ---
auth_manager = SpotifyClientCredentials(client_id=client_id, client_secret=client_secret)
sp = spotipy.Spotify(auth_manager=auth_manager)

artist_ids = [
    'UCS8-ccff2oGbG0Mx7UOuKYw',
    'UCCHDdiih5D__dDi2V3nboPw',
    'UCabtR67_U5O72yRXF7hiI-g',
    'UCgzshmpXAc1T30PHQ3Yw2lw',
    'UCSrA5JaXpR21z_E1FDA03RA',
    'UCPQPjvs4xHLIShnJMXY7GlQ',
    'UCUAa1fv7JsIFuBUWtH0ZaOA',
    'UCZU9T1ceaOgwfLRq7OKFU4Q',
    'UCEuOwB9vSL1oPKGNdONB4ig',
    'UCeMsJJOE6avjyvbqP4Kf24g',
]
"""
    'UCDVnp5x53g5L-kBvQsGOW9w',
    'UCWfAZudbGLAyrrEBaGuslNA',
    'UCrx-X329UKv0Y06VhfpFVvw',
    'UCs6eXM7s8Vl5WcECcRHc2qQ',
    'UC48vpdaG8NDvEGLj11XPPZQ',
    'UCvpredjG93ifbCP1Y77JyFA',
    'UCbb9FsCRA83-6vdgXn2KSyA',
    'UCbcXadHgtd1pWzTAWoI6loA',
    'UCvJLQlVdWiIU0qYmowdKh6g',
    'UCNpdKmV1hHFuKM6DUxGMOBw',
    'UCqLqkJrUJ36z5bm0erMDkjQ',
    'UCqC_GY2ZiENFz2pwL0cSfAw',
    'UCIVROfF-b3Wjpk-G1-_ulFw',
    'UCkhyoTaWKuB-Rdbb6Z3Z5DA',
    'UCiMhD4jzUqG-IgPzUmmytRQ',
    'UCdvlHk5SZWwr9HjUcwtu8ng',
    'UCwolwTPDNMSzHV-HRii4gFQ',
    'UCc4K7bAqpdBP8jh1j9XZAww',
    'UCU3UjFG2RPDP0kBqDTnHOOg',
    'UC7-YMmnc0ppcWmio8t1WdcA',
    'UCB0JSO6d5ysH2Mmqz5I9rIw',
    'UCyOw2FDjfQOFQH7paKxNVvA',
    'UCae-zZkho0dkoQ5Z_NJhpeg',
    'UCBhNW531Uyx1ooaWlcepyvg',
    'UCC6C_vZmq4eMsjnKoNeiP8g',
    'UCFMZHIQMgBXTSxsr86Caazw',
    'UCsQBsZJltmLzlsJNG7HevBg',
    'UCcjWoLUPOkEVhevjI8DB5vQ',
    'UC3lBXcrKFnFAFkfVk5WuKcQ',
    'UCuq1H-HXWoW4JL-hX5bWxzw',
    'UC6Y8dX3XdGsp2akU-SKfreA',
    'UCiKxNv_MHAShqT2lATxG_Wg',
    'UC3WIRbOw46MsbycRc1N5x4g',
    'UCA_jfhUc3VTeMX021QEhBAQ',
    'UCoNPsL8j28yfKRu6e7YUhPA',
    'UCqf-kTp9ERV5T1rPayno7LA',
    'UCY2qt3dw2TQJxvBrDiYGHdQ',
    'UCe3JCin4Gnv9azlWnAs5keg',
    'UCB_Z6rBg3WW3NL4-QimhC2A',
    'UCbulh9WdLtEXiooRcYK7SWw',
    'UCUxNZEOdVy77QiiSTHk8bug',
    'UCi2KNss4Yx73NG0JARSFe0A',
    'UCT9zcQNlyht7fRlcjmflRSA',
    'UCBQZwaNPFfJ1gZ1fLZpAEGw',
    'UC-KTRBl9_6AX10-Y7IKwKdw',
    'UC7IZbkKDd20v6Ni_vFX63rQ',
    'UCq19-LqvG35A-30oyAiPiqA',
    'UC20LoHy2mX0LQODrkUalxVQ',
    'UCWnqnojAgMdN0fQpr_xByJw',
    'UC_kRDKYrUlrbtrSiyu5Tflg',
    'UCLVz1B001PIbq9LliJenV2w',
    'UCbirjI1K3MGu0-Y1gTBNR5w',
    'UCoK2X3nnuG7ug-Kk75r5rJQ',
    'UCP3EX5VKeaG4Oa2qTKPuEFA',
    'UCOJZ1tna8yj8mAEITPkHNCQ',
    'UCfIXdjDQH9Fau7y99_Orpjw',
    'UC69yJGpLNIMk6_ECLwxBZwA',
    'UCvyt-PsgNd1LpxYdxI4vfmA',
    'UCPdGk36EZTJl_vt_y9ZKVSw',
    'UCbFVp7N9HIqmKmB_SsXm-ZA',
    'UCP6HPY4IdCsOC2ap5brcwzw',
    'UCMraDotjZTljynTjcpxLcug',
    'UCmeMTJZRa80cGlMiNpXh64Q',
    """

sql_statements = [
    "-- SQL миграция для вставки данных из YouTube Music\n",
    f"-- Генерация: {datetime.now().isoformat()}\n\n",
    "DO $$\nDECLARE\n"
    "    v_artist_id UUID;\n"
    "    v_album_id UUID;\n"
    "    v_track_id UUID;\n"
    "BEGIN\n"
]

tracks_for_upload = []

for artist_id in artist_ids:
    try:
        # --- Получаем информацию об артисте ---
        artist_info = yt.get_artist(artist_id)
        artist_name = artist_info.get('name').replace("'", "''")
        print(f"=== Обрабатывается артист {artist_name} ===")

        results = sp.search(q=artist_name, type="artist", limit=1)
        spotify_artist = results['artists']['items'][0]
        artist_avatar_url = upload_avatar(spotify_artist['images'][2]['url'], f"artists/{artist_name}_avatar.webp",
                                          bucket_name, minio_client)

        artist_description = (artist_info.get('description', '-') or '').replace("'", "''")
        artist_thumbnails = artist_info.get('thumbnails', [])
        artist_header_url = upload_avatar(artist_thumbnails[-1]['url'], f"artists/{artist_name}_header.webp",
                                          bucket_name, minio_client)

        sql_statements.append(f"-- Артист: {artist_name}\n")
        sql_statements.append(
            f"INSERT INTO artist (artist_name, description, avatar_url, header_url)\n"
            f"VALUES ('{artist_name}', '{artist_description}', '{artist_avatar_url}', '{artist_header_url}')\n"
            f"RETURNING artist_id INTO v_artist_id;\n"
        )

        # --- Получаем все альбомы и синглы ---
        albums_list = []

        # --- Альбомы ---
        albums_info = artist_info.get('albums')
        if albums_info:
            params = albums_info.get('params')
            channelId = albums_info.get('browseId')
            if params:
                full_albums = yt.get_artist_albums(channelId, params=params, limit=None)
                albums_list.extend(full_albums)
            else:
                albums_list.extend(albums_info.get('results', []))

        # --- Синглы ---
        singles_info = artist_info.get('singles')
        if singles_info:
            params = singles_info.get('params')
            channelId = singles_info.get('browseId')
            if params:
                full_singles = yt.get_artist_albums(channelId, params=params, limit=None)
                albums_list.extend(full_singles)
            else:
                albums_list.extend(singles_info.get('results', []))

        min_year, max_year = 3000, 0
        # --- Обработка альбомов и синглов ---
        for alb in albums_list:
            title = alb.get('title', '—').replace("'", "''")
            thumbnails = alb.get('thumbnails', [])
            avatar_url = upload_avatar(thumbnails[-1]['url'], f"albums/{title}.webp", bucket_name, minio_client)
            release_year = alb.get('year')
            alb_type = alb.get('type')

            # --- определяем тип альбома ---
            if alb_type and not str(alb_type).isdigit():
                alb_type = alb_type
            elif release_year and not str(release_year).isdigit():
                t = release_year
                release_year = alb_type
                alb_type = t
            else:
                release_year = alb_type
                alb_type = "Альбом"

            # --- определяем год ---
            if release_year and str(release_year).isdigit():
                release_year = int(release_year)
                min_year = release_year if release_year < min_year else min_year
                max_year = release_year if release_year > max_year else max_year
            else:
                if max_year != 0:
                    if min_year != max_year:
                        release_year = random.randint(min_year, max_year)
                    else:
                        release_year = min_year
                else:
                    release_year = 1999

            sql_statements.append(
                f"\n-- Альбом/сингл: {title}\n"
                f"INSERT INTO album (title, avatar_url, artist_id, description, release_date, type)\n"
                f"VALUES ('{title}', '{avatar_url}', v_artist_id, '', '{release_year}-01-01', '{alb_type}')\n"
                f"RETURNING album_id INTO v_album_id;\n"
            )

            # --- Получаем треки ---
            album_browse_id = alb.get('browseId')
            audio_playlist_id = alb.get('audioPlaylistId')
            if audio_playlist_id:
                tracks = yt.get_playlist(audio_playlist_id, limit=None).get('tracks', [])
            elif album_browse_id:
                tracks = yt.get_album(album_browse_id).get('tracks', [])
            else:
                tracks = []

            sql_statements.append(f"-- Найдено треков: {len(tracks)}\n")

            for track in tracks:
                track_title = track.get('title', '—').replace("'", "''")
                duration_s = int(track.get('duration_seconds', 0))
                videoId = track.get('videoId', '')
                file_url = f"http://localhost:8080/api/v1/track/{videoId}"

                tracks_for_upload.append({
                    "videoId": videoId,
                    "title": track_title,
                })

                sql_statements.append(
                    f"INSERT INTO track (title, duration_s, file_url, description)\n"
                    f"VALUES ('{track_title}', {duration_s}, '{file_url}', '')\n"
                    f"RETURNING track_id INTO v_track_id;\n"
                )

                sql_statements.append(
                    f"INSERT INTO track_artist (track_id, artist_id) VALUES (v_track_id, v_artist_id) ON CONFLICT DO NOTHING;\n"
                )
                sql_statements.append(
                    f"INSERT INTO track_album (track_id, album_id) VALUES (v_track_id, v_album_id) ON CONFLICT DO NOTHING;\n"
                )
    except Exception as e:
        print(f"Ошибка при обработке артиста {artist_id}: {e}")

sql_statements.append("END $$;\n")

# --- Сохраняем SQL ---
with open(output_file, "w", encoding="utf-8") as f:
    f.writelines(sql_statements)

with open("tracks_to_upload.json", "w", encoding="utf-8") as f:
    json.dump(tracks_for_upload, f, ensure_ascii=False, indent=2)

print(f"\nSQL-скрипт успешно создан: {output_file}")