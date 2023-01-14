package apperr_test

import (
	"errors"
	"testing"

	"github.com/percypham/go-apperr"
)

func TestOperationalError(t *testing.T) {
	t.Run("create error with apperr.New", func(t *testing.T) {
		err := apperr.New("error")
		if !apperr.IsOperational(err) {
			t.Error("expected to be operational error")
		}
	})

	t.Run("create apperr.New with info", func(t *testing.T) {
		key, val := "key", "val"
		err := apperr.New("error").With(apperr.Info(key, val))

		info := apperr.OutmostInfo(err, key)

		if info != val {
			t.Errorf("expected '%s', got '%s'", val, info)
		}
	})

	t.Run("wrapped apperr.New", func(t *testing.T) {
		err := apperr.New("err msg")
		wrappedErr := apperr.Wrap(err, "do something")
		if !apperr.IsOperational(wrappedErr) {
			t.Error("expected to be operational error")
		}
	})

	t.Run("error not created with apperr.New", func(t *testing.T) {
		err := errors.New("error")
		if apperr.IsOperational(err) {
			t.Error("expected NOT to be operational error")
		}
	})

	t.Run("wrapped programmer error", func(t *testing.T) {
		err := errors.New("error")
		wrappedErr := apperr.Wrap(err, "do something")
		if apperr.IsOperational(wrappedErr) {
			t.Error("expected NOT to be operational error")
		}
	})
}
