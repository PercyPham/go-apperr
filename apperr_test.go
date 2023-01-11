package apperr_test

import (
	"errors"
	"testing"

	"github.com/percypham/go-apperr"
)

func TestOperationalError(t *testing.T) {
	t.Run("created error with apperr.New", func(t *testing.T) {
		err := apperr.New("error")
		if !apperr.IsOperational(err) {
			t.Error("expected to be operational error, got false")
		}
	})

	t.Run("wrapped created error with apperr.New", func(t *testing.T) {
		err := apperr.New("err msg")
		wrappedErr := apperr.Wrap(err, "do something")
		if !apperr.IsOperational(wrappedErr) {
			t.Error("expected to be operational error, got false")
		}
	})
}

func TestProgrammerError(t *testing.T) {
	t.Run("error not created with apperr.New", func(t *testing.T) {
		err := errors.New("error")
		if apperr.IsOperational(err) {
			t.Error("expected not to be operational error")
		}
	})

	t.Run("wrapped programmer error", func(t *testing.T) {
		err := errors.New("error")
		wrappedErr := apperr.Wrap(err, "do something")
		if apperr.IsOperational(wrappedErr) {
			t.Error("expected not to be operational error")
		}
	})
}
