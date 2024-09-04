package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeats 5 times", func(t *testing.T) {
		repeated := Repeat("a", -1)
		expected := "aaaaa"
	
		assertCorrectMessage(t, repeated, expected)
	})
	t.Run("repeats  twice when interations = 2", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		assertCorrectMessage(t, repeated, expected)
	})
}

func assertCorrectMessage(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("P", 3)
	fmt.Println(repeated)
	// Output: PPP
}