package gosh

import (
	"testing"
)

func TestGosh(t *testing.T) {
	err := Run("date", "cd ..", "pwd", "ls")
	if err != nil {
		t.Error("gosh run error.")
	}
}
