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

	// DefaultServeMux にハンドラが設定される
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		html := `
			<html>
				<head>
					<title>Hello, world!</title>
				</head>
				<body>
					<h1>Welcome!</h1>
					<p>This server is written in Go.</p>
					<p><a href="/hello">/hello</a>
				</body>
			</html>
			`
		res.Write([]byte(html))
	})

	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})

	// 第2引数が nil なら DefaultServeMux が使われる
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
