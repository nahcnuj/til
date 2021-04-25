package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, s := f(5)
	fmt.Println("string "+s+", int ", n)

	fmt.Println(double(3))

	fmt.Println(sum(1, 2, 3))

	f := func(s []string) (string, []string) {
		return s[0], s[1:]
	}

	a := []string{"one", "two", "three"}
	for len(a) > 0 {
		var v string
		v, a = f(a)
		fmt.Println("popped "+v+", remaining: ", a)
	}

	square := func(n int) int {
		return n * n
	}
	b := []int{1, 2, 3}
	fmt.Println(sumf(b, square))

	hello := wrap()
	fmt.Println(hello())

	k := 2
	c := func(n int) int { // クロージャ
		return k * n
	}

	fmt.Println(c(3))

	k = 3
	fmt.Println(c(3))

	fmt.Println(func(n int) int {
		return 2 * n
	}(3))
}

// 複数の戻り値
func f(n int) (int, string) {
	return n, strconv.Itoa(n)
}

// 名前付き戻り値
func double(n int) (d int) {
	d = 2 * n
	return // return は必須
}

// 可変長引数
func sum(a ...int) (s int) { // a はスライスになる
	for _, n := range a {
		s += n
	}
	return
}

// 関数を引数に取る関数
func sumf(a []int, f func(int) int) (s int) {
	for _, n := range a {
		s += f(n)
	}
	return
}

// 関数を返す関数
func wrap() func() string {
	return func() string {
		return "hello"
	}
}
