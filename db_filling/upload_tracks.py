import yt_dlp
import os
import json
from minio import Minio

minio_client = Minio(
    "localhost:8099",
    access_key="miniouser",
    secret_key="miniopassword",
    secure=False
)
bucket_name = "music"
if not minio_client.bucket_exists(bucket_name):
    minio_client.make_bucket(bucket_name)

def download_audio(video_id, tmp_dir="/tmp"):
    url = f"https://www.youtube.com/watch?v={video_id}"
    ydl_opts = {
        'outtmpl': f'{tmp_dir}/{video_id}.webm',
        'format': 'bestaudio/best',
    }

    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=True)
        filename = ydl.prepare_filename(info)
    return filename

with open("data/tracks_to_upload.json", "r", encoding="utf-8") as f:
    tracks = json.load(f)

for t in tracks:
    video_id = t["video_id"]
    title = t["title"]
    print(f"Загрузка трека: {title}")

    try:
        filepath = download_audio(video_id)
        object_name = f"tracks/{os.path.basename(filepath)}"

        with open(filepath, "rb") as data:
            minio_client.put_object(
                bucket_name=bucket_name,
                object_name=object_name,
                data=data,
                length=os.path.getsize(filepath),
                content_type="audio/webm"
            )

        print(f"{title} upload")
    except Exception as e:
        print(f"Ошибка при загрузке {title}: {e}")