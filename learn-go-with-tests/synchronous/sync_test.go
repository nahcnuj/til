package synchronous

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times, leaveing it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("run safely concurrently", func(t *testing.T) {
		wantedCount := 1_000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, c Counter, want int) {
	t.Helper()
	if got := c.Value(); got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
