package apperr_test

import (
	"errors"
	"testing"

	"github.com/percypham/go-apperr"
)

func TestInfo(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		info := apperr.OutmostInfo(nil, "key")

		if info != nil {
			t.Errorf("expected nil, got %v", info)
		}
	})

	t.Run("not set info", func(t *testing.T) {
		err := errors.New("err")
		wrappedErr := apperr.Wrap(err, "do something")

		info := apperr.OutmostInfo(wrappedErr, "non-exist-key")

		if info != nil {
			t.Errorf("expected nil, got %v", info)
		}
	})

	t.Run("exist info", func(t *testing.T) {
		key, val := "key", "val"
		outmostVal := "outmost val"
		err := apperr.New("err").With(apperr.Info(key, val))
		wrappedErr := apperr.Wrap(err, "do something").With(apperr.Info(key, outmostVal))

		info := apperr.OutmostInfo(wrappedErr, key)

		if info != outmostVal {
			t.Errorf("expected '%v', got '%v'", outmostVal, info)
		}
	})

	t.Run("info set in nested err", func(t *testing.T) {
		key, val := "key", "val"
		err := errors.New("err")
		wrappedErr := apperr.Wrap(err, "do something").With(apperr.Info(key, val))
		wrappedErr = apperr.Wrap(wrappedErr, "do something else")

		info := apperr.OutmostInfo(wrappedErr, key)

		if info != val {
			t.Errorf("expected '%v', got '%v'", val, info)
		}
	})

	t.Run("grouped infos", func(t *testing.T) {
		key1, val1 := "key1", "val1"
		key2, val2 := "key2", "val2"
		groupedInfos := apperr.GroupInfos(apperr.Info(key1, val1), apperr.Info(key2, val2))

		err := errors.New("err")
		wrappedErr := apperr.Wrap(err, "do something").With(groupedInfos)
		wrappedErr = apperr.Wrap(wrappedErr, "do something else")

		t.Run("key1", func(t *testing.T) {
			info1 := apperr.OutmostInfo(wrappedErr, key1)
			if info1 != val1 {
				t.Errorf("expected '%v', got '%v'", val1, info1)
			}
		})
		t.Run("key2", func(t *testing.T) {
			info2 := apperr.OutmostInfo(wrappedErr, key2)
			if info2 != val2 {
				t.Errorf("expected '%v', got '%v'", val2, info2)
			}
		})
	})
}
