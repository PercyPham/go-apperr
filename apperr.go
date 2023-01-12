package apperr

import (
	"errors"
	"fmt"
)

func New(text string) *operationalError {
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

	infoMap map[string]any
}

func (e *operationalError) With(infos ...InfoSetter) *operationalError {
	for _, infoSetter := range infos {
		infoSetter.setInfo(e)
	}
	return e
}

func (e *operationalError) setInfo(key string, val any) {
	if e.infoMap == nil {
		e.infoMap = map[string]any{}
	}
	e.infoMap[key] = val
}

func (e *operationalError) Unwrap() error {
	return errOperational
}

func (e *operationalError) Error() string {
	return e.text
}

func (e *operationalError) StackTrace() string {
	if len(e.infoMap) > 0 {
		return fmt.Sprintf("%s | %s\n\t%s", e.text, serializeInfoMap(e.infoMap), e.stackTrace)
	}
	return fmt.Sprintf("%s\n\t%s", e.text, e.stackTrace)
}
