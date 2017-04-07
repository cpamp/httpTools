package httpHelper

import (
	"reflect"
)

// ErrorResponse An error response
type ErrorResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

// SetMessage set the message of the error response
func (er *ErrorResponse) SetMessage(message string) {
	er.Message = message
}

// ZeroedErrorResponse return zeroed error response
func ZeroedErrorResponse() *ErrorResponse {
	return &ErrorResponse{0, "", nil}
}

// IsErrorResponse check if object is an Error response
func IsErrorResponse(i interface{}) bool {
	return reflect.TypeOf(i) == reflect.TypeOf(ZeroedErrorResponse())
}

// StringResponse string response
type StringResponse interface {
	String() string
}

type SafeResponse struct {
	StatusCode int         `json:"statusCode"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
}
