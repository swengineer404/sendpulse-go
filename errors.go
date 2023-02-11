package sendpulse

import "fmt"

const (
	ErrNoData            = 8
	ErrSenderNotFound    = 10
	ErrRecipientNotFound = 11
	ErrEmailNotFound     = 303
)

type APIError struct {
	StatusCode int
	URL        string
	Body       string
	ErrorCode  string `json:"error_code"`
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
