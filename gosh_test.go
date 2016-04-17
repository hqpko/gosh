package gosh

import "testing"

func TestGosh(t *testing.T) {
	err := Run("date")
	if err != nil {
		t.Error("gosh run error.")
	}
}
