// Package sqs provides a SQS client.
package sqs

import (
	"context"
)

// Client defines the AWS SQS client interface.
type Client interface {
	GetFilenamesFromMessages(ctx context.Context) ([]string, error)
}
