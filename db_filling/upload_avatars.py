import requests
from io import BytesIO

def upload_avatar(url, object_name, bucket_name, minio_client):
    if not url:
        return ''

    response = requests.get(url, timeout=60)
    if response.status_code != 200:
        response = requests.get(url, timeout=60)
        if response.status_code != 200:
            return ''

    data = BytesIO(response.content)

    minio_client.put_object(
        bucket_name=bucket_name,
        object_name=object_name,
        data=data,
        length=len(data.getbuffer()),
        content_type='image/webp',
    )

    return f"http://localhost:8099/avatars/{object_name}"
