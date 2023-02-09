package errors

import (
	"errors"
	"fmt"
)

// ErrInvalidArgument is returned when an argument is invalid.
type ErrInvalidArgument struct {
	// message is the error message.
	message string
}

// NewErrInvalidArgument creates a new ErrInvalidArgument.
func NewErrInvalidArgument(message string) *ErrInvalidArgument {
	return &ErrInvalidArgument{
		message: message,
	}
}

// Error implements the error interface.
func (e *ErrInvalidArgument) Error() string {
	return e.message
}

// IsErrInvalidArgument checks if the given error is an ErrInvalidArgument.
func IsErrInvalidArgument(err error) bool {
	var x *ErrInvalidArgument
	return errors.As(err, &x)
}

// ErrNotFound is returned when a resource is not found.
type ErrNotFound struct {
	// message is the error message.
	message string
}

// NewErrNotFound creates a new ErrNotFound.
func NewErrNotFound(message string) *ErrNotFound {
	return &ErrNotFound{
		message: message,
	}
}

// Error implements the error interface.
func (e *ErrNotFound) Error() string {
	return e.message
}

// IsErrNotFound checks if the given error is an ErrNotFound.
func IsErrNotFound(err error) bool {
	var x *ErrNotFound
	return errors.As(err, &x)
}

// ErrInternal is returned when an internal error occurs.
type ErrInternal struct {
	// message is the error message.
	message string
}

// NewErrInternal creates a new ErrInternal.
func NewErrInternal(message string) *ErrInternal {
	return &ErrInternal{
		message: message,
	}
}

// Error implements the error interface.
func (e *ErrInternal) Error() string {
	return e.message
}

// ErrUnauthorized is returned when a user is unauthorized.
type ErrUnauthorized struct {
	// message is the error message.
	message string
}

// NewErrUnauthorized creates a new ErrUnauthorized.
func NewErrUnauthorized(message string) *ErrUnauthorized {
	return &ErrUnauthorized{
		message: message,
	}
}

// Error implements the error interface.
func (e *ErrUnauthorized) Error() string {
	return e.message
}

// IsErrUnauthorized checks if the given error is an ErrUnauthorized.
func IsErrUnauthorized(err error) bool {
	var x *ErrUnauthorized
	return errors.As(err, &x)
}

// ErrForbidden is returned when a user is forbidden.
type ErrForbidden struct {
	// message is the error message.
	message string
}

// NewErrForbidden creates a new ErrForbidden.
func NewErrForbidden(message string) *ErrForbidden {
	return &ErrForbidden{
		message: message,
	}
}

// Error implements the error interface.
func (e *ErrForbidden) Error() string {
	return e.message
}

// IsErrForbidden checks if the given error is an ErrForbidden.
func IsErrForbidden(err error) bool {
	var x *ErrForbidden
	return errors.As(err, &x)
}

// ErrEntityNotFound is returned when an entity is not found.
type ErrEntityNotFound struct {
	actor  string
	source string
}

// NewErrEntityNotFound creates a new ErrEntityNotFound.
func NewErrEntityNotFound(actor string, source string) *ErrEntityNotFound {
	return &ErrEntityNotFound{
		actor:  actor,
		source: source,
	}
}

// Error implements the error interface.
func (e *ErrEntityNotFound) Error() string {
	return fmt.Sprintf("Entity not found (%s): %s", e.actor, e.source)
}

// IsErrEntityNotFound checks if the given error is an ErrEntityNotFound.
func IsErrEntityNotFound(err error) bool {
	var x *ErrEntityNotFound
	return errors.As(err, &x)
}
