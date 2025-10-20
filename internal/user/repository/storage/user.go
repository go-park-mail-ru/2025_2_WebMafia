package storage

import (
	"bytes"
	"context"
	"fmt"
)

func (s *Storage) UploadAvatar(ctx context.Context, objectName string, data []byte, contentType string) error {
	reader := bytes.NewReader(data)

	if err := s.client.Upload(ctx, s.bucket, objectName, reader, int64(len(data)), contentType); err != nil {
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
