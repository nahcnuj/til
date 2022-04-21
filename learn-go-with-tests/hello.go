package main

import "fmt"

const japanese = "Japanese"
const french = "French"
const englishHelloPrefix = "Hello, "
const japaneseHelloPrefix = "こんにちは、"
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
