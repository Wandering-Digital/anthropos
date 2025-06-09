package response

import (
	"encoding/json"
	"net/http"

	"github.com/Wandering-Digital/anthropos/internal/customerror"
	"github.com/Wandering-Digital/anthropos/internal/paginator"
)

func write(w http.ResponseWriter, statusCode int, resp *Response) error {
	if resp == nil {
		w.WriteHeader(statusCode)

		return nil
	}

	if statusCode >= http.StatusInternalServerError && resp.Error != nil {
		resp.Message = resp.Error.Message
		resp.Error = nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}

	return nil
}

type Response struct {
	Message    string                `json:"message,omitempty"`
	Data       interface{}           `json:"data,omitempty"`
	Error      *customerror.APIError `json:"error,omitempty"`
	Pagination *paginator.Pagination `json:"pagination,omitempty"`
}

// WithData writes data with success status code.
func WithData(w http.ResponseWriter, statusCode int, resp Response) error {
	return write(w, statusCode, &resp)
}

// WithPaginatedData writes paginated data to the Response.
func WithPaginatedData(w http.ResponseWriter, statusCode int, data interface{}, pagination *paginator.Pagination) error {
	return write(w, statusCode, &Response{Data: data, Pagination: pagination})
}

func WithSuccessNoContent(w http.ResponseWriter) error {
	return write(w, http.StatusNoContent, nil)
}

// WithError write Response with custom error.
func WithError(w http.ResponseWriter, err *customerror.APIError) error {
	return write(w, err.Code, &Response{Error: err})
}
