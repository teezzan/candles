// Package sqs provides a SQS client.
package sqs

import (
	"context"
)
//go:generate moq -rm -out client_mock.go . Client

// Client defines the AWS SQS client interface.
type Client interface {
	GetFilenamesFromMessages(ctx context.Context) ([]string, error)
	DeleteMessages(ctx context.Context, messageHandles []string) error
}
