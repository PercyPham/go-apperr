package apperr_test

import (
	"errors"
	"testing"

	"github.com/percypham/go-apperr"
)

func TestWrapError(t *testing.T) {
	t.Run("wrap nil", func(t *testing.T) {
		appErr := apperr.Wrap(nil, "do something")
		if appErr != nil {
			t.Errorf("expected nil AppError, got %v", appErr)
		}
	})

	t.Run("wrap normal error", func(t *testing.T) {
		err := errors.New("err msg")
		appErr := apperr.Wrap(err, "do %s", "something")

		if appErr == nil {
			t.Fatalf("expected AppError, got nil")
		}

		expectedLogMsg := "do something: err msg"
		actualLogMsg := appErr.Error()

		if actualLogMsg != expectedLogMsg {
			t.Errorf("expected '%s', got '%s'", expectedLogMsg, actualLogMsg)
		}
	})

	t.Run("wrap appErr", func(t *testing.T) {
		err := errors.New("err msg")
		nestedAppErr := apperr.Wrap(err, "do something")
		appErr := apperr.Wrap(nestedAppErr, "do other thing")

		expectedLogMsg := "do other thing: do something: err msg"
		actualLogMsg := appErr.Error()

		if actualLogMsg != expectedLogMsg {
			t.Errorf("expected '%s', got '%s'", expectedLogMsg, actualLogMsg)
		}
	})
}
