package sqs

import (
	"context"
	"fmt"

	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/teezzan/ohlc/internal/config"
	"github.com/tidwall/gjson"

	"go.uber.org/zap"
)

var _ Client = (*DefaultClient)(nil)

// DefaultClient implements the default AWS S3 client.
type DefaultClient struct {
	logger    *zap.Logger
	sqsClient *sqs.Client
	queueURL  string
}

// NewClient initializes a new default AWS SQS client.
func NewClient(
	ctx context.Context,
	logger *zap.Logger,
	conf config.SQSConfig,

) (*DefaultClient, error) {
	sdkConfig, err := s3Config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := sqs.NewFromConfig(sdkConfig)

	gQInput := &sqs.GetQueueUrlInput{
		QueueName: &conf.Queue,
	}

	// Get URL of queue
	urlResult, err := client.GetQueueUrl(ctx, gQInput)
	if err != nil {
		return nil, err
	}

	return &DefaultClient{
		logger:    logger,
		sqsClient: client,
		queueURL:  *urlResult.QueueUrl,
	}, nil
}

// GetFilenamesFromMessages listens for messages on the queue.
func (c *DefaultClient) GetFilenamesFromMessages(ctx context.Context) ([]string, error) {
	// Receive messages from queue by polling
	result, err := c.sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: &c.queueURL,
	})
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, m := range result.Messages {
		if m.Body != nil {
			keys, err := extractKeyFromMessage(*m.Body)
			if err != nil {
				continue
			}
			filenames = append(filenames, keys...)
		}
	}
	return filenames, nil
}

func extractKeyFromMessage(m string) ([]string, error) {
	var keys []string
	n := gjson.Get(m, "Records.#")
	for i := 0; i < int(n.Int()); i++ {
		s3 := gjson.Get(m, fmt.Sprintf("Records.%d.s3", i))
		if !s3.Exists() {
			continue
		}
		object := gjson.Get(m, fmt.Sprintf("Records.%d.s3.object", i))
		if !object.Exists() {
			continue
		}
		key := gjson.Get(m, fmt.Sprintf("Records.%d.s3.object.key", i))
		if !key.Exists() {
			continue
		}
		keys = append(keys, key.String())
	}
	return keys, nil
}
