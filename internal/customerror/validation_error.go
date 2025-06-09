package customerror

import "encoding/json"

type ValidationError struct {
	validationErrors
}

type validationErrors map[string][]interface{}

func NewValidationError() *ValidationError {
	return &ValidationError{}
}

func (ve ValidationError) Add(key string, value interface{}) {
	ve.validationErrors[key] = append(ve.validationErrors[key], value)
}

func (ve ValidationError) IsNil() bool {
	return len(ve.validationErrors) == 0
}

func (ve ValidationError) Error() string {
	// return ve.validationErrors.Error()

	b, _ := json.Marshal(ve.validationErrors)
	return string(b)
}
