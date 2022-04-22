package countdown

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	// spy records how a dependency is used
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough or too much calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
