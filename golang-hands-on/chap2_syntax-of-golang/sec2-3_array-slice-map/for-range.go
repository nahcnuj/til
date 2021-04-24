package main

import "fmt"

func main() {
	var a = [3]int{4, 5, 6}
	for i, x := range a {
		fmt.Println(i, x)
	}
}
