// Package awsS3 provides an S3 clients.
package awsS3

import "context"

// Client defines the AWS S3 client interface.
type Client interface {
	ListBuckets(ctx context.Context) error
	GeneratePresignedURL(ctx context.Context, key string) (string, error)
}
