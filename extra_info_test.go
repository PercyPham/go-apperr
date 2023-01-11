package apperr_test

import (
	"errors"
	"testing"

	"github.com/percypham/go-apperr"
)

func TestExtraInfo(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		extraInfo := apperr.OutmostExtraInfo(nil, "key")

		if extraInfo != nil {
			t.Errorf("expected nil, got %v", extraInfo)
		}
	})

	t.Run("not set extra info", func(t *testing.T) {
		err := errors.New("err")
		wrappedErr := apperr.Wrap(err, "do something")

		extraInfo := apperr.OutmostExtraInfo(wrappedErr, "non-exist-key")

		if extraInfo != nil {
			t.Errorf("expected nil, got %v", extraInfo)
		}
	})

	t.Run("exist extra info", func(t *testing.T) {
		key, val := "key", "val"
		err := errors.New("err")
		wrappedErr := apperr.Wrap(err, "do something").With(apperr.ExtraInfo(key, val))

		extraInfo := apperr.OutmostExtraInfo(wrappedErr, key)

		if extraInfo != val {
			t.Errorf("expected '%v', got '%v'", val, extraInfo)
		}
	})

	t.Run("extra info set in nested err", func(t *testing.T) {
		key, val := "key", "val"
		err := errors.New("err")
		wrappedErr := apperr.Wrap(err, "do something").With(apperr.ExtraInfo(key, val))
		wrappedErr = apperr.Wrap(wrappedErr, "do something else")

		extraInfo := apperr.OutmostExtraInfo(wrappedErr, key)

		if extraInfo != val {
			t.Errorf("expected '%v', got '%v'", val, extraInfo)
		}
	})
}
