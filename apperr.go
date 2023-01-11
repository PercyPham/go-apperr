package apperr

import (
	"errors"
)

func New(text string) error {
	return &operationalError{
		text:       text,
		stackTrace: getCallerStackTrace(),
	}
}

func IsOperational(err error) bool {
	return errors.Is(err, errOperational)
}

var errOperational = errors.New("operational")

type operationalError struct {
	text       string
	stackTrace string
}

func (e *operationalError) Unwrap() error {
	return errOperational
}

func (e *operationalError) Error() string {
	return e.text
}

func (e *operationalError) StackTrace() string {
	return e.text + "\n\t" + e.stackTrace
}
