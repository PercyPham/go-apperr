package apperr_test

import (
	"testing"

	"github.com/percypham/go-apperr"
)

func TestProgrammerError(t *testing.T) {
	t.Run("create err with apperr.NewProgrammerError", func(t *testing.T) {
		err := apperr.NewProgrammerError("programmer err")
		if apperr.IsOperational(err) {
			t.Error("expected NOT to be operational error")
		}
	})

	t.Run("create apperr.NewProgrammerError with info", func(t *testing.T) {
		key, val := "key", "val"
		err := apperr.NewProgrammerError("error").With(apperr.Info(key, val))

		info := apperr.OutmostInfo(err, key)

		if info != val {
			t.Errorf("expected '%s', got '%s'", val, info)
		}
	})

	t.Run("wrapped apperr.NewProgrammerError", func(t *testing.T) {
		err := apperr.NewProgrammerError("programmer err")
		wrappedErr := apperr.Wrap(err, "do something")

		if apperr.IsOperational(wrappedErr) {
			t.Error("expected NOT to be operational error")
		}
	})
}
