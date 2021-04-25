package main

import "fmt"

func main() {
	m := map[string]int{
		"key": 42,
		"a":   334,
	}
	m["sum"] = m["key"] + m["a"]
	fmt.Println(m)

	delete(m, "a")
	fmt.Println(m)

	delete(m, "b") // no error

	for k, v := range m {
		fmt.Println(k+": ", v)
	}
}
