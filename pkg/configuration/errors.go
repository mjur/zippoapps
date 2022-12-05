package configuration

import (
	"net/http"
)

// BadRequestError occurs when an object is invalid.
type BadRequestError struct {
	text string
	Code int
}

// NewBadRequestError creates a new not found error.
func NewBadRequestError(text string) *BadRequestError {
	return &BadRequestError{text: text, Code: http.StatusBadRequest}
}

// Error returns the error as a string.
func (e *BadRequestError) Error() string {
	return e.text
}

// UnauthorizedError occurs when the user does not have permission to execute.
type UnauthorizedError struct {
	text string
	Code int
}

// NewBadRequestError creates a new not found error.
func NewUnauthorizedError(text string) *UnauthorizedError {
	return &UnauthorizedError{text: text, Code: http.StatusUnauthorized}
}

// Error returns the error as a string.
func (e *UnauthorizedError) Error() string {
	return e.text
}

// NotFoundError occurs when an object wasn't found.
type NotFoundError struct {
	text string
	Code int
}

// NewNotFoundError creates a new not found error.
func NewNotFoundError(text string) *NotFoundError {
	return &NotFoundError{text: text, Code: http.StatusNotFound}
}

// Error returns the error as a string.
func (e *NotFoundError) Error() string {
	return e.text
}

// ServiceError.
type ServiceError struct {
	text string
	Code int
}

// NewNotFoundError creates a new not found error.
func NewServiceError(text string) *ServiceError {
	return &ServiceError{text: text, Code: http.StatusInternalServerError}
}

// Error returns the error as a string.
func (e *ServiceError) Error() string {
	return e.text
}

// ServiceUnavailibleError.
type ServiceUnavailibleError struct {
	text string
	Code int
}

// NewNotFoundError creates a new not found error.
func NewServiceUnavailibleError(text string) *ServiceUnavailibleError {
	return &ServiceUnavailibleError{text: text, Code: http.StatusServiceUnavailable}
}

// Error returns the error as a string.
func (e *ServiceUnavailibleError) Error() string {
	return e.text
}
