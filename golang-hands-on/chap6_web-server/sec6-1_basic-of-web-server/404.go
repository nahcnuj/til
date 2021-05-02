package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	if err := http.ListenAndServe("0.0.0.0:8080", http.NotFoundHandler()); err != nil {
		panic(err)
	}
}
