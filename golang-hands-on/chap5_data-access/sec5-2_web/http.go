package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	r, err := http.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
