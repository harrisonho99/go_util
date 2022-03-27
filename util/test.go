package util

import (
	"testing"
)

func CheckTestError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
