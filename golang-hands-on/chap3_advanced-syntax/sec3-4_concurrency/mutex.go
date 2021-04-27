package main

import (
	"fmt"
	"sync"
	"time"
)

type ShareData struct {
	count int
	mutex sync.Mutex
}

func main() {
	d := ShareData{count: 0}

	print := func(prefix string) {
		fmt.Println(prefix, d.count)
	}

	main := func() {
		for i := 0; i < 10; i++ {
			d.mutex.Lock()
			d.count++
			print("*main")
			d.mutex.Unlock()
		}
	}

	hello := func() {
		for i := 0; i < 10; i++ {
			d.mutex.Lock()
			d.count++
			print("hello")
			d.mutex.Unlock()
		}
	}

	go main()
	go hello()
	time.Sleep(1 * time.Second)
}
