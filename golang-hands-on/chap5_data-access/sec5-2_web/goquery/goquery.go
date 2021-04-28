package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	r, err := http.Get("https://www3.nhk.or.jp/news/")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	d, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		panic(err)
	}

	d.Find(".content--list dd a em").Each(func(n int, s *goquery.Selection) {
		fmt.Println("- " + s.Text())
	})
}
