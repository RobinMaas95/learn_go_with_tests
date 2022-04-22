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

func TestAdd(t *testing.T) {
	dict := Dictionary{}
	dict.Add("test", definition)
	assertDefinition(t, dict, "test", definition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	assertStrings(t, got, definition)
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
