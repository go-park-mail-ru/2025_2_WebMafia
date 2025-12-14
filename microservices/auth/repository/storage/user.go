package storage

import (
	"context"
	"fmt"
	"io"
	"spotify/pkg/minio"
)

func (s *Storage) UploadAvatar(ctx context.Context, file io.Reader, size int64, contentType string) (string, error) {
	obj := minio.ObjectInfo{
		Bucket:      s.bucket,
		Reader:      file,
		Size:        size,
		ContentType: contentType,
	}

	objectName, err := s.client.Upload(ctx, obj)
	if err != nil {
		return "", fmt.Errorf("upload avatar: %w", ErrUploadFailed)
	}
	return objectName, nil
}

func (s *Storage) DeleteAvatar(ctx context.Context, objectName string) error {
	if err := s.client.Remove(ctx, s.bucket, objectName); err != nil {
		return fmt.Errorf("delete avatar: %w", ErrDeleteFailed)
	}
	return nil
}
