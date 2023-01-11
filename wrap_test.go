package apperr_test

import (
	"errors"
	"testing"

	"github.com/percypham/go-apperr"
)

func TestWrapNil(t *testing.T) {
	wrappedErr := apperr.Wrap(nil, "do something")
	if wrappedErr != nil {
		t.Errorf("expected nil, got '%v'", wrappedErr)
	}
}

func TestWrapProgrammerError(t *testing.T) {
	t.Run("wrap 1 layer", func(t *testing.T) {
		err := errors.New("programmer error")
		wrappedErr := apperr.Wrap(err, "do something")

		if wrappedErr == nil {
			t.Fatal("expected err, got nil")
		}

		expectedLogMsg := "do something: programmer error"
		actualLogMsg := wrappedErr.Error()

		if actualLogMsg != expectedLogMsg {
			t.Errorf("expected '%s', got '%s'", expectedLogMsg, actualLogMsg)
		}
	})

	t.Run("wrap multiple layers", func(t *testing.T) {
		err := errors.New("programmer error")
		wrappedErr := apperr.Wrap(err, "do something 1")
		wrappedErr = apperr.Wrap(wrappedErr, "do something 2")

		if wrappedErr == nil {
			t.Fatal("expected err, got nil")
		}

		expectedLogMsg := "do something 2: do something 1: programmer error"
		actualLogMsg := wrappedErr.Error()

		if actualLogMsg != expectedLogMsg {
			t.Errorf("expected '%s', got '%s'", expectedLogMsg, actualLogMsg)
		}
	})
}
