package blogrenderer

import (
	"blogposts"
	"fmt"
	"io"
)

func Render(w io.Writer, post blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>\n", post.Title)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "<p>%s</p>\n", post.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "Tags: <ul>")
	if err != nil {
		return err
	}

	for _, tag := range post.Tags {
		fmt.Fprintf(w, "<li>%s</li>", tag)
	}

	_, err = fmt.Fprintf(w, "</ul>")
	return err
}
