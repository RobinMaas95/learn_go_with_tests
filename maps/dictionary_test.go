package main

import "testing"

var definition = "This is just a test"
func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": definition}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, got, definition)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("bla")
		assertError(t, err, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
