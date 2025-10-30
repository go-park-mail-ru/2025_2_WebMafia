package minio

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	UseSSL    bool
}

type ObjectInfo struct {
	Bucket      string
	ObjectName  string
	Reader      io.Reader
	Size        int64
	ContentType string
}

type Client struct {
	minioClient *minio.Client
}

func New(cfg Config) (*Client, error) {
	mClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("init minio client: %w", err)
	}

	return &Client{minioClient: mClient}, nil
}

func (c *Client) EnsureBucket(ctx context.Context, bucket string) error {
	exists, err := c.minioClient.BucketExists(ctx, bucket)
	if err != nil {
		return fmt.Errorf("check bucket: %w", err)
	}
	if !exists {
		if err := c.minioClient.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		log.Printf("Bucket %s создан", bucket)
	}
	return nil
}

func (c *Client) Upload(ctx context.Context, obj ObjectInfo) (string, error) {
	if err := c.EnsureBucket(ctx, obj.Bucket); err != nil {
		return "", err
	}

	ext := strings.TrimPrefix(obj.ContentType, "image/")
	objectName := fmt.Sprintf("%s.%s", uuid.New().String(), ext)

	_, err := c.minioClient.PutObject(ctx, obj.Bucket, objectName, obj.Reader, obj.Size, minio.PutObjectOptions{
		ContentType: obj.ContentType,
	})

	if err != nil {
		return "", fmt.Errorf("upload object: %w", err)
	}
	return objectName, nil
}

func (c *Client) Remove(ctx context.Context, bucket, objectName string) error {
	if err := c.minioClient.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{}); err != nil {
		return fmt.Errorf("remove object: %w", err)
	}
	return nil
}

func (c *Client) GetFileURL(ctx context.Context, bucket, objectName string) (string, error) {
	params := make(url.Values)
	u, err := c.minioClient.PresignedGetObject(ctx, bucket, objectName, time.Hour, params)
	if err != nil {
		return "", fmt.Errorf("presign url: %w", err)
	}
	return u.String(), nil
}
