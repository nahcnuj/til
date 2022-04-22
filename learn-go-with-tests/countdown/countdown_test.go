package countdown

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	// spy records how a dependency is used
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const (
	write = "write"
	sleep = "sleep"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
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
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write, // 3
			sleep,
			write, // 2
			sleep,
			write, // 1
			sleep,
			write, // Go!
		}
		got := spySleepPrinter.Calls

		if !reflect.DeepEqual(want, got) {
			t.Errorf("wanted calls %v got %v", want, got)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
