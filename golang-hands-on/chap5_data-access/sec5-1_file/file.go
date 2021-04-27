package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		fmt.Println(string(b))
	}

	_, err = os.ReadFile("not-found-file.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	os.WriteFile("output.txt", []byte("こんにちは、world"), os.ModePerm)
}
