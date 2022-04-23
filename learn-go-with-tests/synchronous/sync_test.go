package synchronous

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times, leaveing it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
}

func assertCounter(t testing.TB, c Counter, want int) {
	t.Helper()
	if got := c.Value(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
