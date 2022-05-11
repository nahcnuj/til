package app

import (
	"sync"
	"testing"
)

func TestStoreConcurrently(t *testing.T) {
	wanted_count := 1_000
	player := "Pepper"
	store := NewInMemoryPlayerStore()

	var wg sync.WaitGroup
	wg.Add(wanted_count)

	for i := 0; i < wanted_count; i++ {
		go func() {
			store.RecordWin(player)
			wg.Done()
		}()
	}
	wg.Wait()

	got := store.GetPlayerScore(player)
	if got != wanted_count {
		t.Errorf("got %d wins, but want %d wins", got, wanted_count)
	}
}
