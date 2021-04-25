package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type number int

func (num number) isPrime() bool {
	n := int(num)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func (num number) primeFactors() []int {
	f := []int{}
	n := int(num)
	p := 2
	for n > p {
		if n%p == 0 {
			f = append(f, p)
			n /= p
		} else {
			if p == 2 {
				p++
			} else {
				p += 2
			}
		}
	}
	return append(f, n)
}

func main() {
	s := input("type a number")
	if x, err := strconv.Atoi(s); err == nil {
		n := number(x)
		fmt.Printf("Is %d a prime? => %t\n", n, n.isPrime())
		fmt.Println(n.primeFactors())

		n = 2*n + 1
		fmt.Printf("Is %d a prime? => %t\n", n, n.isPrime())
		fmt.Println(n.primeFactors())
	}
}

func input(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}
