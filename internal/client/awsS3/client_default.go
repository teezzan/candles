package awsS3

import (
	"context"
	"fmt"
	"time"

	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/teezzan/ohlc/internal/config"

	"go.uber.org/zap"
)

var _ Client = (*DefaultClient)(nil)

// DefaultClient implements the default AWS S3 client.
type DefaultClient struct {
	logger               *zap.Logger
	s3Client             *s3.Client
	presignClient        *s3.PresignClient
	bucketName           string
	region               string
	presignURLExpiryTime int
}

// NewClient initializes a new default AWS S3 client.
func NewClient(
	ctx context.Context,
	logger *zap.Logger,
	conf config.S3Config,

) (*DefaultClient, error) {
	sdkConfig, err := s3Config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	presignClient := s3.NewPresignClient(s3Client)
	return &DefaultClient{
		logger:               logger,
		s3Client:             s3Client,
		presignClient:        presignClient,
		bucketName:           conf.Bucket,
		region:               conf.Region,
		presignURLExpiryTime: conf.PresignURLExpiryTime * 60,
	}, nil
}

// ListBuckets lists all the buckets.
func (c *DefaultClient) ListBuckets(ctx context.Context) error {
	result, err := c.s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return err
	}
	for _, bucket := range result.Buckets {
		fmt.Printf("\t%v\n", *bucket.Name)
	}
	return nil
}

// GeneratePresignedURL generates a presigned URL for the given bucket and key.
func (c *DefaultClient) GeneratePresignedURL(ctx context.Context, key string) (string, error) {
	req, err := c.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(int64(c.presignURLExpiryTime) * int64(time.Second))
	})

	if err != nil {
		return "", err
	}
	return req.URL, nil
}
