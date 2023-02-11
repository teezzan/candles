// Package s3 provides an S3 clients.
package s3

import "context"

//go:generate moq -rm -out client_mock.go . Client

// Client defines the AWS S3 client interface.
type Client interface {
	ListBuckets(ctx context.Context) error
	GeneratePresignedURL(ctx context.Context, key string) (string, error)
	DownloadLargeObject(ctx context.Context, objectKey string) ([]byte, error)
}
