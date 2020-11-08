package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Raj", "")
		want := "Hello, Raj!"

		assertMessage(t, got, want)
	})

	t.Run("say `Hello World!` on empty name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertMessage(t, got, want)
	})

	t.Run("say in French", func(t *testing.T) {
		got := Hello("Raj", "French")
		want := "Bonjour, Raj!"

		assertMessage(t, got, want)
	})
}
