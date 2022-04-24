package blogrenderer

import (
	"blogposts"
	"fmt"
	"io"
)

func Render(w io.Writer, post blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", post.Title)
	return err
}
