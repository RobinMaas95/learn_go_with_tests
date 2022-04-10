package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMesssage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Robin", "English")
		want := "Hello, Robin"
		assertCorrectMesssage(t, got, want)

	})
	t.Run("saying'Hello World' when a empty string is passed", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMesssage(t, got, want)

	})
	t.Run("saying the greeting in Spanish", func(t *testing.T) {
		got := Hello("Robin", "Spanish")
		want := "Hola, Robin"
		assertCorrectMesssage(t, got, want)
	})
	t.Run("saying the greeting in German", func(t *testing.T) {
		got := Hello("Robin", "German")
		want := "Hallo, Robin"
		assertCorrectMesssage(t, got, want)
	})
	t.Run("Using English as default when a unknown langauge was passed", func(t *testing.T) {
		got := Hello("Robin", "Italian")
		want := "Hello, Robin"
		assertCorrectMesssage(t, got, want)
	})
}
