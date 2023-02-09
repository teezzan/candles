package httputil

import (
	"net/http"

	E "github.com/teezzan/ohlc/internal/errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorResponse is the error response.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponseFromError creates a new ErrorResponse from an error.
func NewErrorResponseFromError(err error, code int) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}

// HandlerWrapperWithLoggerFunc is a function that wraps a http.HandlerFunc with error.
type HandlerWrapperWithLoggerFunc func(func(c *gin.Context) error) gin.HandlerFunc

func NewHandlerWrapper(logger *zap.Logger) HandlerWrapperWithLoggerFunc {
	return func(fn func(c *gin.Context) error) gin.HandlerFunc {
		return func(c *gin.Context) {
			err := fn(c)
			if err == nil {
				return
			}

			if err != nil {
				logger.Error("error in handler", zap.Error(err))

				if E.IsErrInvalidArgument(err) {
					BadRequest(c, err)
				} else if E.IsErrUnauthorized(err) {
					Unauthorized(c, err)
				} else if E.IsErrForbidden(err) {
					Forbidden(c, err)
				} else if E.IsErrNotFound(err) {
					NotFound(c, err)
				} else {
					InternalServerError(c)
				}
			}
		}
	}
}

// OK responds with a 200 OK status code and JSON payload if provided.
func OK(c *gin.Context, data interface{}) error {
	return respond(c, http.StatusOK, data)
}

// BadRequest responds with a 400 Bad Request status code and JSON payload if provided.
func BadRequest(c *gin.Context, err error) error {
	errData := NewErrorResponseFromError(err, http.StatusBadRequest)
	return respond(c, http.StatusBadRequest, errData)
}

// Unauthorized responds with a 401 Unauthorized status code and JSON payload if provided.
func Unauthorized(c *gin.Context, err error) error {
	errData := NewErrorResponseFromError(err, http.StatusUnauthorized)
	return respond(c, http.StatusUnauthorized, errData)
}

// Forbidden responds with a 403 Forbidden status code and JSON payload if provided.
func Forbidden(c *gin.Context, err error) error {
	errData := NewErrorResponseFromError(err, http.StatusForbidden)
	return respond(c, http.StatusForbidden, errData)
}

// NotFound responds with a 404 Not Found status code and JSON payload if provided.
func NotFound(c *gin.Context, err error) error {
	errData := NewErrorResponseFromError(err, http.StatusNotFound)
	return respond(c, http.StatusNotFound, errData)
}

// NoContent responds with a 204 No Content status code.
func NoContent(c *gin.Context) error {
	return respond(c, http.StatusNoContent, nil)
}

// InternalServerError responds with a 500 Internal Server Error status code and JSON payload if provided.
func InternalServerError(c *gin.Context) error {
	return respond(c, http.StatusInternalServerError, "Internal Server Error")
}

func respond(c *gin.Context, statusCode int, data interface{}) error {
	c.AbortWithStatusJSON(statusCode, data)
	return nil
}
