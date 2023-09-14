package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Isaac")
	got := buffer.String()
	want := "Hello, Isaac"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
