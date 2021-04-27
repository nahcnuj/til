package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	printFile := func(f *os.File) {
		r := bufio.NewReaderSize(f, 4096)
		for i := 1; true; i++ {
			b, _, err := r.ReadLine()
			if err != nil {
				if err.Error() == "EOF" {
					break
				} else {
					panic(err)
				}
			}
			fmt.Println(i, ": ", string(b))
		}
	}

	f, err := os.OpenFile("data.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	printFile(f)
}
