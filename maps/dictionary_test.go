package main

import "testing"

func TestSearch(t *testing.T) {
	key := "test"
	value := "This is just a test"
	dictionary := map[string]string{key: value}
	got := Search(dictionary, key)
	want := value

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, key)
	}
}
