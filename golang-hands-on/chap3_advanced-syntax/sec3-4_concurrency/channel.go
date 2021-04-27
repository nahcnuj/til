package main

import (
	"fmt"
	"strconv"
	"time"
)

func total(n int, c chan int) {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	c <- s
}

func total2(c chan int) {
	n := <-c

	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	c <- s
}

func printAndWait(s string, n uint) {
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func first(n uint, c chan string) {
	for i := 0; i < 10; i++ {
		s := "first-" + strconv.Itoa(i)
		printAndWait(s, n)
		c <- s
	}
}

func second(n uint, c chan string) {
	for i := 0; i < 10; i++ {
		printAndWait("second: ["+<-c+"]", n)
	}
}

func total3(recv chan int, send chan int) {
	n := <-recv
	fmt.Println("n = ", n)

	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	send <- s
}

func count(n int, t uint, c chan int) {
	for i := 1; i <= n; i++ {
		c <- i
		time.Sleep(time.Duration(t) * time.Millisecond)
	}
}

func main() {
	{
		c := make(chan int)
		go total(1000, c)
		go total(100, c)
		go total(10, c)
		x, y, z := <-c, <-c, <-c
		fmt.Println(x, y, z)

		go total2(c)
		c <- 10000
		fmt.Println(<-c)
	}

	{
		c := make(chan string)
		go first(10, c)
		second(10, c)
		fmt.Println()
	}

	{
		send := make(chan int)
		recv := make(chan int)

		go total3(send, recv)

		send <- 100
		fmt.Println(<-recv)
	}

	{
		n := []int{3, 5, 10}
		t := []uint{100, 75, 50}
		c := []chan int{make(chan int), make(chan int), make(chan int)}
		for i := 0; i < len(c); i++ {
			go count(n[i], t[i], c[i])
		}

		for i := 0; i < n[0]+n[1]+n[2]; i++ {
			select {
			case r := <-c[0]:
				fmt.Println("*   1st ", r)
			case r := <-c[1]:
				fmt.Println("**  2nd ", r)
			case r := <-c[2]:
				fmt.Println("*** 3rd ", r)
			}
		}
		fmt.Println("====")
	}
}
