package main

import "fmt"

func main() {
	n, m := 123, 42
	p, q := &n, &m
	pp, qq := &p, &q

	fmt.Printf("pp addr=%p value=%d\n", *pp, **pp)
	fmt.Printf("qq addr=%p value=%d\n", *qq, **qq)

	swap(&p, &q)

	fmt.Printf("pp addr=%p value=%d\n", *pp, **pp)
	fmt.Printf("qq addr=%p value=%d\n", *qq, **qq)
}

func swap(a **int, b **int) {
	t := *a
	*a = *b
	*b = t
}
