package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	{
		tmpl, err := template.New("").Parse("{{.Text}}")
		if err != nil {
			panic(err)
		}

		http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			if err := tmpl.Execute(res, struct{ Text string }{Text: "Hello"}); err != nil {
				panic(err)
			}
		})
	}

	{
		tmpl, err := template.ParseFiles("templates/hello.html", "templates/good-morning.html")
		if err != nil {
			panic(err)
		}

		http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
			data := struct {
				Title   string
				Message string
			}{
				Title:   "Hi",
				Message: "Welcome",
			}
			if err := tmpl.Execute(res, data); err != nil { // execute the first template, templates/hello.html
				panic(err)
			}
		})
	}

	// ANY /hello/:template/
	{
		tmpl, err := template.ParseGlob("templates/*")
		if err != nil {
			panic(err)
		}

		ptn := "/hello/"
		http.HandleFunc(ptn, func(res http.ResponseWriter, req *http.Request) {
			// deal path parameters
			path := strings.TrimPrefix(req.URL.EscapedPath(), ptn)
			segs := strings.Split(path, "/")

			// deal query string parameters
			if err := req.ParseForm(); err != nil {
				panic(err)
			}

			t := tmpl.Lookup(segs[0] + ".html")
			if t == nil {
				http.NotFound(res, req)
				return
			}

			data := struct {
				Title    string
				Message  string
				Segments []string
				Queries  url.Values
			}{
				Title:    "Hi",
				Message:  "Welcome",
				Segments: segs,
				Queries:  req.Form,
			}
			if err := t.Execute(res, data); err != nil {
				fmt.Fprintln(os.Stderr, err)
				http.NotFound(res, req)
			}
		})
	}

	http.ListenAndServe(":8080", nil)
}
