package storage

import (
	"spotify/pkg/minio"
)

type Storage struct {
	client *minio.Client
	bucket string
}

func NewStorage(client *minio.Client, bucket string) *Storage {
	return &Storage{
		client: client,
		bucket: bucket,
	}
}
