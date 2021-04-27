package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walk(f string, ds []string) {
	h := append(ds, f)
	p := filepath.Join(h...)
	s, err := os.Stat(p)
	if err != nil {
		panic(err)
	}

	fmt.Print(strings.Repeat("  ", len(ds)), s.Name())
	if s.IsDir() {
		fmt.Println("/")
	} else {
		fmt.Printf(" ( %d )\n", s.Size())
		return
	}

	fs, err := os.ReadDir(p)
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		walk(f.Name(), h)
	}
}

func printDirTree(p string) {
	p, err := filepath.Abs(p)
	if err != nil {
		panic(err)
	}
	walk(p, []string{})
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], " {path}")
		return
	}

	p := os.Args[1]
	printDirTree(p)
}
