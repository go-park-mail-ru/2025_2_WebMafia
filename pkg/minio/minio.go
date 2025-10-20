package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
	"net/url"
	"time"
)

type Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
}

type Client struct {
	minioClient *minio.Client
	bucket      string
}

func New(cfg Config) (*Client, error) {
	mClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("init minio client: %w", err)
	}

	ctx := context.Background()
	exists, err := mClient.BucketExists(ctx, cfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("check bucket: %w", err)
	}

	if !exists {
		err = mClient.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("create bucket: %w", err)
		}
		log.Printf("Bucket %s создан", cfg.Bucket)
	}

	return &Client{minioClient: mClient, bucket: cfg.Bucket}, nil
}

func (c *Client) Upload(ctx context.Context, bucket, objectName string, reader io.Reader, size int64, contentType string) error {
	_, err := c.minioClient.PutObject(ctx, bucket, objectName, reader, size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("upload object: %w", err)
	}
	return nil
}

func (c *Client) Remove(ctx context.Context, bucket, objectName string) error {
	err := c.minioClient.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
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
