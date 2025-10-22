package storage

import (
	"bytes"
	"context"
	"fmt"
	"spotify/pkg/minio"
)

func (s *Storage) UploadAvatar(ctx context.Context, objectName string, data []byte, contentType string) error {
	obj := minio.ObjectInfo{
		Bucket:      s.bucket,
		ObjectName:  objectName,
		Reader:      bytes.NewReader(data),
		Size:        int64(len(data)),
		ContentType: contentType,
	}

	if err := s.client.Upload(ctx, obj); err != nil {
		return fmt.Errorf("upload avatar: %w", ErrUploadFailed)
	}
	return nil
}

func (s *Storage) DeleteAvatar(ctx context.Context, objectName string) error {
	if err := s.client.Remove(ctx, s.bucket, objectName); err != nil {
		return fmt.Errorf("delete avatar: %w", ErrDeleteFailed)
	}
	return nil
}

func (s *Storage) GetAvatarURL(ctx context.Context, objectName string) (string, error) {
	url, err := s.client.GetFileURL(ctx, s.bucket, objectName)
	if err != nil {
		return "", fmt.Errorf("get avatar url: %w", ErrURLFailed)
	}
	return url, nil
}
