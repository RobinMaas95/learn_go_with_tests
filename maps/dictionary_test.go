package main

import "testing"

func TestSearch(t *testing.T) {
	value := "This is just a test"
	dictionary := Dictionary{"test": value}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		assertStrings(t, got, value)
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
