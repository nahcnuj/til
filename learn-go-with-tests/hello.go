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
	if lang == japanese {
		return japaneseHelloPrefix + name
	}
	if lang == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
