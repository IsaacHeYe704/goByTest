package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Isaac", "Spanish")
		want := "Hola, Isaac"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Isaac", "French")
		want := "Bonjour, Isaac"
		assertCorrectMessage(t, got, want)
	})
}
func assertCorrectMessage(t testing.TB, got string, want string) {
	//this function is a helper so if it fals the test suit will mark the ccde that calls this function and not this lines
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
