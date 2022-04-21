package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, want, got string) {
		t.Helper() // tells the test suite that this assertion method is a helper
		if got != want {
			t.Errorf("expected %q but got %q", want, got)
		}
	}

	t.Run("repeat 'a' 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, expected, repeated)
	})

	t.Run("repeat 'x' 10 times", func(t *testing.T) {
		repeated := Repeat("x", 10)
		expected := "xxxxxxxxxx"

		assertCorrectMessage(t, expected, repeated)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
