package http

import (
	"net/http"

	"github.com/mjur/zippo/pkg/configuration"
	"github.com/pkg/errors"
)

// ResponseError represent the reseponse error struct.
type ResponseError struct {
	Message string `json:"message,omitempty"`
}

func GetErrorResponse(err error) (int, *ResponseError) {
	if err == nil {
		return http.StatusOK, nil
	}

	var status int

	var respError ResponseError

	switch errors.Cause(err).(type) {
	case *configuration.BadRequestError:
		status = http.StatusBadRequest
		respError.Message = err.Error()
	case *configuration.NotFoundError:
		status = http.StatusNotFound
		respError.Message = err.Error()
	case *configuration.UnauthorizedError:
		status = http.StatusUnauthorized
		respError.Message = "Unauthorized"
	case *configuration.ServiceUnavailibleError:
		status = http.StatusServiceUnavailable
		respError.Message = "Service unavailable"
	default:
		status = http.StatusInternalServerError
		respError.Message = "Internal server error"
	}

	return status, &respError
}
