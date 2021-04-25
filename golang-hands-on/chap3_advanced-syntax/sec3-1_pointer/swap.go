package main

import "fmt"

func main() {
	a := 42
	b := 334

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
