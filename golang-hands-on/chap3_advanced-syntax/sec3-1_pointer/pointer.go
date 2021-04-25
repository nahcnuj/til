package main

import "fmt"

func main() {
	n := 123
	p := &n
	fmt.Println("pointer ", p)
	fmt.Println("value: ", *p)
}
