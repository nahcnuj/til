package synchronous

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times, leaveing it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d, want %d", counter.Value(), 3)
		}
	})
}
