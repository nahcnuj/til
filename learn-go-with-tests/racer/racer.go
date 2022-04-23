package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func ping(url string) chan struct{} {
	// need to make not to get nil, or it blocks forever because Go cannot send to it
	ch := make(chan struct{}) // struct{} allocates nothing, is smaller than bool
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
