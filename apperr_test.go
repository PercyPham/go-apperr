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
		appErr := apperr.Wrapf(err, "do %s", "something")

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

func TestRootError(t *testing.T) {
	t.Run("root err of nil", func(t *testing.T) {
		appErr := apperr.RootError(nil)

		if appErr != nil {
			t.Errorf("expected nil, got %v", appErr)
		}
	})

	t.Run("root err of normal err", func(t *testing.T) {
		err := errors.New("root err")

		rootErr := apperr.RootError(err)

		if rootErr != err {
			t.Errorf("expected '%v', got '%v'", err, rootErr)
		}
	})

	t.Run("root err of appErr", func(t *testing.T) {
		err := errors.New("root err")
		appErr := apperr.Wrap(err, "do something")

		rootErr := apperr.RootError(appErr)

		if rootErr != err {
			t.Errorf("expected '%v', got '%v'", err, rootErr)
		}
	})
}
