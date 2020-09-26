package main

import (
	"testing"
)

func TestMessage(t *testing.T) {

	var result = Message("Hello")

	if result != "<h1> Hello </h1>" {
		t.Errorf("Result is invalid: returned: %v, expected: %v", result, "<h1> Hello </h1>")
	}
}
