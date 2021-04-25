package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	initialize(&a)
	fmt.Println(a)
}

func initialize(a *[]int) {
	for i := 0; i < len(*a); i++ {
		(*a)[i] = 0
	}
}
