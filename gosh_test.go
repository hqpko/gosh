package gosh

import "testing"

func TestGosh(t *testing.T) {
	EchoTime = true
	EchoIn = true
	EchoOut = true
	err := Run("date")
	if err != nil {
		t.Error("gosh run error.")
	}
}
