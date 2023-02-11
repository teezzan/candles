// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package repository

import (
	"context"
	"github.com/teezzan/ohlc/internal/controller/ohlc/data"
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			GetDataPointsFunc: func(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error) {
//				panic("mock out the GetDataPoints method")
//			},
//			InsertDataPointsFunc: func(ctx context.Context, rows []data.OHLCEntity) error {
//				panic("mock out the InsertDataPoints method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// GetDataPointsFunc mocks the GetDataPoints method.
	GetDataPointsFunc func(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error)

	// InsertDataPointsFunc mocks the InsertDataPoints method.
	InsertDataPointsFunc func(ctx context.Context, rows []data.OHLCEntity) error

	// calls tracks calls to the methods.
	calls struct {
		// GetDataPoints holds details about calls to the GetDataPoints method.
		GetDataPoints []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Payload is the payload argument value.
			Payload data.GetOHLCRequest
		}
		// InsertDataPoints holds details about calls to the InsertDataPoints method.
		InsertDataPoints []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Rows is the rows argument value.
			Rows []data.OHLCEntity
		}
	}
	lockGetDataPoints    sync.RWMutex
	lockInsertDataPoints sync.RWMutex
}

// GetDataPoints calls GetDataPointsFunc.
func (mock *RepositoryMock) GetDataPoints(ctx context.Context, payload data.GetOHLCRequest) ([]data.OHLCEntity, error) {
	if mock.GetDataPointsFunc == nil {
		panic("RepositoryMock.GetDataPointsFunc: method is nil but Repository.GetDataPoints was just called")
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
//	len(mockedRepository.GetDataPointsCalls())
func (mock *RepositoryMock) GetDataPointsCalls() []struct {
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

// InsertDataPoints calls InsertDataPointsFunc.
func (mock *RepositoryMock) InsertDataPoints(ctx context.Context, rows []data.OHLCEntity) error {
	if mock.InsertDataPointsFunc == nil {
		panic("RepositoryMock.InsertDataPointsFunc: method is nil but Repository.InsertDataPoints was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Rows []data.OHLCEntity
	}{
		Ctx:  ctx,
		Rows: rows,
	}
	mock.lockInsertDataPoints.Lock()
	mock.calls.InsertDataPoints = append(mock.calls.InsertDataPoints, callInfo)
	mock.lockInsertDataPoints.Unlock()
	return mock.InsertDataPointsFunc(ctx, rows)
}

// InsertDataPointsCalls gets all the calls that were made to InsertDataPoints.
// Check the length with:
//
//	len(mockedRepository.InsertDataPointsCalls())
func (mock *RepositoryMock) InsertDataPointsCalls() []struct {
	Ctx  context.Context
	Rows []data.OHLCEntity
} {
	var calls []struct {
		Ctx  context.Context
		Rows []data.OHLCEntity
	}
	mock.lockInsertDataPoints.RLock()
	calls = mock.calls.InsertDataPoints
	mock.lockInsertDataPoints.RUnlock()
	return calls
}