package s3

import (
	"context"
	"fmt"
	"time"

	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/teezzan/candles/internal/config"

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
	sdkConfig, err := s3Config.LoadDefaultConfig(ctx, func(lo *s3Config.LoadOptions) error {
		lo.Region = conf.Region
		return nil
	})
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
		presignURLExpiryTime: conf.PresignURLExpiryTime * 3600,
	}, nil
}

// ListBuckets lists all the available buckets in the S3 client
// It prints the name of each bucket to the standard output.
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

// GeneratePresignedURL returns a presigned URL for the provided object key in the specified bucket.
// The URL will expire after the time specified in the `presignURLExpiryTime` field.
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

// DownloadLargeObject downloads a large object from an Amazon S3 bucket. It downloads the object
// in parts of 10 MiBs and returns the object as a slice of bytes. The object key is passed as a
// parameter to the function. The function returns an error if the download fails.
func (c *DefaultClient) DownloadLargeObject(ctx context.Context, objectKey string) ([]byte, error) {
	var partMiBs int64 = 10
	downloader := manager.NewDownloader(c.s3Client, func(d *manager.Downloader) {
		d.PartSize = partMiBs * 1024 * 1024
	})
	buffer := manager.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(c.bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), err
}
