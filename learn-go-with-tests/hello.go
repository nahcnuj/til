package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == "Japanese" {
		return "こんにちは、" + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
