// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package sqs

import (
	"context"
	"sync"
)

// Ensure, that ClientMock does implement Client.
// If this is not the case, regenerate this file with moq.
var _ Client = &ClientMock{}

// ClientMock is a mock implementation of Client.
//
//	func TestSomethingThatUsesClient(t *testing.T) {
//
//		// make and configure a mocked Client
//		mockedClient := &ClientMock{
//			DeleteMessagesFunc: func(ctx context.Context, messageHandles []string) error {
//				panic("mock out the DeleteMessages method")
//			},
//			GetFilenamesFromMessagesFunc: func(ctx context.Context) ([]string, error) {
//				panic("mock out the GetFilenamesFromMessages method")
//			},
//		}
//
//		// use mockedClient in code that requires Client
//		// and then make assertions.
//
//	}
type ClientMock struct {
	// DeleteMessagesFunc mocks the DeleteMessages method.
	DeleteMessagesFunc func(ctx context.Context, messageHandles []string) error

	// GetFilenamesFromMessagesFunc mocks the GetFilenamesFromMessages method.
	GetFilenamesFromMessagesFunc func(ctx context.Context) ([]string, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteMessages holds details about calls to the DeleteMessages method.
		DeleteMessages []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// MessageHandles is the messageHandles argument value.
			MessageHandles []string
		}
		// GetFilenamesFromMessages holds details about calls to the GetFilenamesFromMessages method.
		GetFilenamesFromMessages []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockDeleteMessages           sync.RWMutex
	lockGetFilenamesFromMessages sync.RWMutex
}

// DeleteMessages calls DeleteMessagesFunc.
func (mock *ClientMock) DeleteMessages(ctx context.Context, messageHandles []string) error {
	if mock.DeleteMessagesFunc == nil {
		panic("ClientMock.DeleteMessagesFunc: method is nil but Client.DeleteMessages was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		MessageHandles []string
	}{
		Ctx:            ctx,
		MessageHandles: messageHandles,
	}
	mock.lockDeleteMessages.Lock()
	mock.calls.DeleteMessages = append(mock.calls.DeleteMessages, callInfo)
	mock.lockDeleteMessages.Unlock()
	return mock.DeleteMessagesFunc(ctx, messageHandles)
}

// DeleteMessagesCalls gets all the calls that were made to DeleteMessages.
// Check the length with:
//
//	len(mockedClient.DeleteMessagesCalls())
func (mock *ClientMock) DeleteMessagesCalls() []struct {
	Ctx            context.Context
	MessageHandles []string
} {
	var calls []struct {
		Ctx            context.Context
		MessageHandles []string
	}
	mock.lockDeleteMessages.RLock()
	calls = mock.calls.DeleteMessages
	mock.lockDeleteMessages.RUnlock()
	return calls
}

// GetFilenamesFromMessages calls GetFilenamesFromMessagesFunc.
func (mock *ClientMock) GetFilenamesFromMessages(ctx context.Context) ([]string, error) {
	if mock.GetFilenamesFromMessagesFunc == nil {
		panic("ClientMock.GetFilenamesFromMessagesFunc: method is nil but Client.GetFilenamesFromMessages was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetFilenamesFromMessages.Lock()
	mock.calls.GetFilenamesFromMessages = append(mock.calls.GetFilenamesFromMessages, callInfo)
	mock.lockGetFilenamesFromMessages.Unlock()
	return mock.GetFilenamesFromMessagesFunc(ctx)
}

// GetFilenamesFromMessagesCalls gets all the calls that were made to GetFilenamesFromMessages.
// Check the length with:
//
//	len(mockedClient.GetFilenamesFromMessagesCalls())
func (mock *ClientMock) GetFilenamesFromMessagesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetFilenamesFromMessages.RLock()
	calls = mock.calls.GetFilenamesFromMessages
	mock.lockGetFilenamesFromMessages.RUnlock()
	return calls
}
