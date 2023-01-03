package apperr_test

import (
	"errors"
	"testing"

	"github.com/percypham/go-apperr"
)

func TestWrapErrorWithOutOption(t *testing.T) {
	// TODO: add tests
}

func TestWrapErrorWithOption(t *testing.T) {
	t.Run("wrap nil", func(t *testing.T) {
		appErr := apperr.Wrap(nil, "do something").With(apperr.Code(1))
		if appErr != nil {
			t.Errorf("expected nil, got %v", appErr)
		}
	})

	t.Run("wrap error", func(t *testing.T) {
		err := errors.New("err msg")
		appErr := apperr.Wrap(err, "do something").With(apperr.Code(1))

		if appErr == nil {
			t.Fatalf("expected appErr, got nil")
		}

		// TODO: add tests for options
	})
}
