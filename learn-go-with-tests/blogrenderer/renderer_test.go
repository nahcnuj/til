package blogrenderer_test

import (
	"blogposts"
	"blogrenderer"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var aPost = blogposts.Post{
		Title:       "Hello, world",
		Description: "This is a description.",
		Tags:        []string{"go", "TDD"},
		Body:        "This is a blog post.",
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("convert a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var aPost = blogposts.Post{
		Title:       "Hello, world",
		Description: "This is a description.",
		Tags:        []string{"go", "TDD"},
		Body:        "This is a blog post.",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogrenderer.Render(io.Discard, aPost)
	}
}
