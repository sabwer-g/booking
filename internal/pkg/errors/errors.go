package errors

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	requestValidationError = "400"
	invalidJSONCode        = "409"
	internalErrorCode      = "500"
)

const (
	requestValidationMsg = "Validation error"
	invalidJSONMsg       = "Invalid message JSON format"
	internalErrorMsg     = "Internal Server Error"
)

type HTTPError struct {
	Code        string      `json:"code"`
	Message     string      `json:"message"`
	Description string      `json:"description,omitempty"`
	Extra       interface{} `json:"extra,omitempty"`
}

func extractErrorCaller() string {
	_, file, line, _ := runtime.Caller(2)
	paths := strings.Split(file, "/")
	paths = paths[len(paths)-3:]
	return fmt.Sprintf("%s:%d", strings.Join(paths, "/"), line)
}

func (e *HTTPError) Error() string {
	if e.Description != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Description)
	}

	return e.Message
}

func NewRequestValidationError(description string) *HTTPError {
	return &HTTPError{
		Code:        requestValidationError,
		Message:     requestValidationMsg,
		Description: description,
		Extra:       extractErrorCaller(),
	}
}

func NewInvalidJSONError(description string) *HTTPError {
	return &HTTPError{
		Code:        invalidJSONCode,
		Message:     invalidJSONMsg,
		Description: description,
		Extra:       extractErrorCaller(),
	}
}

func NewInternalError(description string) *HTTPError {
	return &HTTPError{
		Code:        internalErrorCode,
		Message:     internalErrorMsg,
		Description: description,
		Extra:       extractErrorCaller(),
	}
}
