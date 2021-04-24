package main

import "fmt"

func main() {
	var a = [...]int{4, 5, 6}
	for _, x := range a {
		fmt.Println(x)
	}
}
