package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)

	a = push(a, 4)
	fmt.Println(a)

	a = pop(a)
	fmt.Println(a)

	a = unshift(a, 5)
	fmt.Println(a)

	a = shift(a)
	fmt.Println(a)
}

func push(s []int, v int) []int {
	return append(s, v)
}

func pop(s []int) []int {
	return s[:len(s)-1]
}

func unshift(s []int, v int) []int {
	return append([]int{v}, s...)
}

func shift(s []int) []int {
	return s[1:]
}

func insert(s []int, v int, i int) []int {
	return append(s[:i], append([]int{v}, s[i+1:]...)...)
}

func remove(s []int, i int) []int {
	return append(s[:i-1], s[i+1:]...)
}
