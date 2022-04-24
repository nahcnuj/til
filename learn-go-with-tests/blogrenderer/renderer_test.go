package blogrenderer_test

import (
	"blogrenderer"
	"blogposts"
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var aPost = blogposts.Post{
		Title:       "Hello, world",
		Description: "This is a description.",
		Tags:        []string{"go", "TDD"},
		Body:        "This is a blog post.",
	}

	t.Run("convert a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello, world</h1>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
