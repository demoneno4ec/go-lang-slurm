package main

import (
	"strings"
	"testing"
)

func Test_foo(t *testing.T) {
	if err := foo(); !strings.Contains(err.Error(), panicMessage) {
		t.Errorf("foo() = %v, want error with %s", err, panicMessage)
	}
}
