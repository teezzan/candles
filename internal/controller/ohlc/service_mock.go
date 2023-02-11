// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ohlc

import (
	"context"
	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
	"sync"
)

// Ensure, that ServiceMock does implement Service.
// If this is not the case, regenerate this file with moq.
var _ Service = &ServiceMock{}

// ServiceMock is a mock implementation of Service.
//
//	func TestSomethingThatUsesService(t *testing.T) {
//
//		// make and configure a mocked Service
//		mockedService := &ServiceMock{
//			CreateDataPointsFunc: func(ctx context.Context, dataPoints [][]string) error {
//				panic("mock out the CreateDataPoints method")
//			},
//			DownloadAndProcessCSVFunc: func(ctx context.Context, filename string) error {
//				panic("mock out the DownloadAndProcessCSV method")
//			},
//			GeneratePreSignedURLFunc: func(ctx context.Context) (*data.GeneratePresignedURLResponse, error) {
//				panic("mock out the GeneratePreSignedURL method")
//			},
//			GetAndProcessSQSMessageFunc: func(ctx context.Context) error {
//				panic("mock out the GetAndProcessSQSMessage method")
//			},
//			GetDataPointsFunc: func(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, *int, error) {
//				panic("mock out the GetDataPoints method")
//			},
//		}
//
//		// use mockedService in code that requires Service
//		// and then make assertions.
//
//	}
type ServiceMock struct {
	// CreateDataPointsFunc mocks the CreateDataPoints method.
	CreateDataPointsFunc func(ctx context.Context, dataPoints [][]string) error

	// DownloadAndProcessCSVFunc mocks the DownloadAndProcessCSV method.
	DownloadAndProcessCSVFunc func(ctx context.Context, filename string) error

	// GeneratePreSignedURLFunc mocks the GeneratePreSignedURL method.
	GeneratePreSignedURLFunc func(ctx context.Context) (*data.GeneratePresignedURLResponse, error)

	// GetAndProcessSQSMessageFunc mocks the GetAndProcessSQSMessage method.
	GetAndProcessSQSMessageFunc func(ctx context.Context) error

	// GetDataPointsFunc mocks the GetDataPoints method.
	GetDataPointsFunc func(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, *int, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateDataPoints holds details about calls to the CreateDataPoints method.
		CreateDataPoints []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DataPoints is the dataPoints argument value.
			DataPoints [][]string
		}
		// DownloadAndProcessCSV holds details about calls to the DownloadAndProcessCSV method.
		DownloadAndProcessCSV []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filename is the filename argument value.
			Filename string
		}
		// GeneratePreSignedURL holds details about calls to the GeneratePreSignedURL method.
		GeneratePreSignedURL []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAndProcessSQSMessage holds details about calls to the GetAndProcessSQSMessage method.
		GetAndProcessSQSMessage []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetDataPoints holds details about calls to the GetDataPoints method.
		GetDataPoints []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Payload is the payload argument value.
			Payload data.GetOHLCRequest
		}
	}
	lockCreateDataPoints        sync.RWMutex
	lockDownloadAndProcessCSV   sync.RWMutex
	lockGeneratePreSignedURL    sync.RWMutex
	lockGetAndProcessSQSMessage sync.RWMutex
	lockGetDataPoints           sync.RWMutex
}

// CreateDataPoints calls CreateDataPointsFunc.
func (mock *ServiceMock) CreateDataPoints(ctx context.Context, dataPoints [][]string) error {
	if mock.CreateDataPointsFunc == nil {
		panic("ServiceMock.CreateDataPointsFunc: method is nil but Service.CreateDataPoints was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		DataPoints [][]string
	}{
		Ctx:        ctx,
		DataPoints: dataPoints,
	}
	mock.lockCreateDataPoints.Lock()
	mock.calls.CreateDataPoints = append(mock.calls.CreateDataPoints, callInfo)
	mock.lockCreateDataPoints.Unlock()
	return mock.CreateDataPointsFunc(ctx, dataPoints)
}

// CreateDataPointsCalls gets all the calls that were made to CreateDataPoints.
// Check the length with:
//
//	len(mockedService.CreateDataPointsCalls())
func (mock *ServiceMock) CreateDataPointsCalls() []struct {
	Ctx        context.Context
	DataPoints [][]string
} {
	var calls []struct {
		Ctx        context.Context
		DataPoints [][]string
	}
	mock.lockCreateDataPoints.RLock()
	calls = mock.calls.CreateDataPoints
	mock.lockCreateDataPoints.RUnlock()
	return calls
}

// DownloadAndProcessCSV calls DownloadAndProcessCSVFunc.
func (mock *ServiceMock) DownloadAndProcessCSV(ctx context.Context, filename string) error {
	if mock.DownloadAndProcessCSVFunc == nil {
		panic("ServiceMock.DownloadAndProcessCSVFunc: method is nil but Service.DownloadAndProcessCSV was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Filename string
	}{
		Ctx:      ctx,
		Filename: filename,
	}
	mock.lockDownloadAndProcessCSV.Lock()
	mock.calls.DownloadAndProcessCSV = append(mock.calls.DownloadAndProcessCSV, callInfo)
	mock.lockDownloadAndProcessCSV.Unlock()
	return mock.DownloadAndProcessCSVFunc(ctx, filename)
}

// DownloadAndProcessCSVCalls gets all the calls that were made to DownloadAndProcessCSV.
// Check the length with:
//
//	len(mockedService.DownloadAndProcessCSVCalls())
func (mock *ServiceMock) DownloadAndProcessCSVCalls() []struct {
	Ctx      context.Context
	Filename string
} {
	var calls []struct {
		Ctx      context.Context
		Filename string
	}
	mock.lockDownloadAndProcessCSV.RLock()
	calls = mock.calls.DownloadAndProcessCSV
	mock.lockDownloadAndProcessCSV.RUnlock()
	return calls
}

// GeneratePreSignedURL calls GeneratePreSignedURLFunc.
func (mock *ServiceMock) GeneratePreSignedURL(ctx context.Context) (*data.GeneratePresignedURLResponse, error) {
	if mock.GeneratePreSignedURLFunc == nil {
		panic("ServiceMock.GeneratePreSignedURLFunc: method is nil but Service.GeneratePreSignedURL was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGeneratePreSignedURL.Lock()
	mock.calls.GeneratePreSignedURL = append(mock.calls.GeneratePreSignedURL, callInfo)
	mock.lockGeneratePreSignedURL.Unlock()
	return mock.GeneratePreSignedURLFunc(ctx)
}

// GeneratePreSignedURLCalls gets all the calls that were made to GeneratePreSignedURL.
// Check the length with:
//
//	len(mockedService.GeneratePreSignedURLCalls())
func (mock *ServiceMock) GeneratePreSignedURLCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGeneratePreSignedURL.RLock()
	calls = mock.calls.GeneratePreSignedURL
	mock.lockGeneratePreSignedURL.RUnlock()
	return calls
}

// GetAndProcessSQSMessage calls GetAndProcessSQSMessageFunc.
func (mock *ServiceMock) GetAndProcessSQSMessage(ctx context.Context) error {
	if mock.GetAndProcessSQSMessageFunc == nil {
		panic("ServiceMock.GetAndProcessSQSMessageFunc: method is nil but Service.GetAndProcessSQSMessage was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAndProcessSQSMessage.Lock()
	mock.calls.GetAndProcessSQSMessage = append(mock.calls.GetAndProcessSQSMessage, callInfo)
	mock.lockGetAndProcessSQSMessage.Unlock()
	return mock.GetAndProcessSQSMessageFunc(ctx)
}

// GetAndProcessSQSMessageCalls gets all the calls that were made to GetAndProcessSQSMessage.
// Check the length with:
//
//	len(mockedService.GetAndProcessSQSMessageCalls())
func (mock *ServiceMock) GetAndProcessSQSMessageCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAndProcessSQSMessage.RLock()
	calls = mock.calls.GetAndProcessSQSMessage
	mock.lockGetAndProcessSQSMessage.RUnlock()
	return calls
}

// GetDataPoints calls GetDataPointsFunc.
func (mock *ServiceMock) GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, *int, error) {
	if mock.GetDataPointsFunc == nil {
		panic("ServiceMock.GetDataPointsFunc: method is nil but Service.GetDataPoints was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Payload data.GetOHLCRequest
	}{
		Ctx:     ctx,
		Payload: payload,
	}
	mock.lockGetDataPoints.Lock()
	mock.calls.GetDataPoints = append(mock.calls.GetDataPoints, callInfo)
	mock.lockGetDataPoints.Unlock()
	return mock.GetDataPointsFunc(ctx, payload)
}

// GetDataPointsCalls gets all the calls that were made to GetDataPoints.
// Check the length with:
//
//	len(mockedService.GetDataPointsCalls())
func (mock *ServiceMock) GetDataPointsCalls() []struct {
	Ctx     context.Context
	Payload data.GetOHLCRequest
} {
	var calls []struct {
		Ctx     context.Context
		Payload data.GetOHLCRequest
	}
	mock.lockGetDataPoints.RLock()
	calls = mock.calls.GetDataPoints
	mock.lockGetDataPoints.RUnlock()
	return calls
}