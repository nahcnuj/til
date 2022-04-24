package blogrenderer_test

import (
	"blogposts"
	"blogrenderer"
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
		want := `<h1>Hello, world</h1><p>This is a description.</p>Tags: <ul><li>go</li><li>TDD</li></ul>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
