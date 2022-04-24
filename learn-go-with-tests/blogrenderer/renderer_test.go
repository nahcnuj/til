package blogrenderer_test

import (
	"blogposts"
	"blogrenderer"
	"bytes"
	"io"
	"strings"
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

	t.Run("render an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogposts.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("render a Markdown body", func(t *testing.T) {
		buf := bytes.Buffer{}
		post := blogposts.Post{Body: `# Heading

Rendered successfully!`}

		if err := postRenderer.Render(&buf, post); err != nil {
			t.Fatal(err)
		}

		got := buf.String()

		wants := []string{`Heading`, `Rendered successfully!`}
		for _, want := range wants {
			if !strings.Contains(got, want) {
				t.Errorf("expected to contain %q", want)
			}
		}

		if strings.Contains(got, `# Heading`) {
			t.Errorf("expected that headings in the body are converted in HTML")
		}
	})
}

func BenchmarkRender(b *testing.B) {
	var aPost = blogposts.Post{
		Title:       "Hello, world",
		Description: "This is a description.",
		Tags:        []string{"go", "TDD"},
		Body:        "This is a blog post.",
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
