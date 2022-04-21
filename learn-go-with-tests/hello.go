package main

import "fmt"

const japanese = "Japanese"
const englishHelloPrefix = "Hello, "
const japaneseHelloPrefix = "こんにちは、"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	if lang == japanese {
		return japaneseHelloPrefix + name
	}
	if lang == "French" {
		return "Bonjour, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
