package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	msg := "start"

	print := func(label string, n int) {
		fmt.Println(label, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
	hello := func(n int) {
		for i := 0; i < 10; i++ {
			msg += " h" + strconv.Itoa(i)
			print("hello", n)
		}
	}
	main := func(n int) {
		for i := 0; i < 5; i++ {
			msg += " m" + strconv.Itoa(i)
			print("*main", n)
		}
	}

	go hello(60)
	main(100)
}
