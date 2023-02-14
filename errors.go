package sendpulse

import (
	"errors"
	"fmt"
)

const (
	ErrorCodeNoData            = 8
	ErrorCodeSenderNotFound    = 10
	ErrorCodeRecipientNotFound = 11
	ErrorCodeEmailNotFound     = 303
)

type APIError struct {
	StatusCode int
	URL        string
	Body       string
	ErrorCode  int    `json:"error_code"`
	Message    string `json:"message"`
	Err        error
}

func NewAPIError(statusCode int, url, body string, err error) *APIError {
	return &APIError{
		StatusCode: statusCode,
		URL:        url,
		Body:       body,
		Err:        err,
	}
}

func (e *APIError) Error() (s string) {
	s = fmt.Sprintf("api error [%s](%d): %s", e.URL, e.StatusCode, e.Body)
	if e.Err != nil {
		s += fmt.Sprintf(" (%s)", e.Err)
	}

	return s
}

func ErrIsCode(err error, code int) bool {
	var e *APIError
	if !errors.As(err, &e) {
		return false
	}

	return e.ErrorCode == code
}
