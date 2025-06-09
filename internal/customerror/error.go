package customerror

import "encoding/json"

type APIError struct {
	Code          int              `json:"code"`
	Message       string           `json:"message,omitempty"`
	OriginalError error            `json:"error,omitempty"`
	Errors        *ValidationError `json:"errors,omitempty"`
}

func (err *APIError) Error() string {
	b, _ := json.Marshal(err)
	return string(b)
}

// NewError is the factory function for generating a custom error payload.
func NewError(code int, message string, err error) *APIError {
	return &APIError{
		Code:          code,
		Message:       message,
		OriginalError: err,
	}
}

// NewErrors is the factory function for generating multi custom error payload.
func NewErrors(code int, message string, errs *ValidationError) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
		Errors:  errs,
	}
}
