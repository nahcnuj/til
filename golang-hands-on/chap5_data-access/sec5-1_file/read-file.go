package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	printFile := func(f *os.File) {
		b, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
	}

	f, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	printFile(f)
}
