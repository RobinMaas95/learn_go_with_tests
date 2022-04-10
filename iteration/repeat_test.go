package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func TestRepeatStrings(t *testing.T) {
	repeated := RepeatUsingStrings("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatUsingStrings("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func ExampleRepeatStrings() {
	repeated := RepeatUsingStrings("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}
