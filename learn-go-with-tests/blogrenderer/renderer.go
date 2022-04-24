package blogrenderer

import (
	"blogposts"
	"embed"
	"html/template"
	"io"
)

//go:embed "templates/*"
var blogTemplates embed.FS

func Render(w io.Writer, post blogposts.Post) error {
	tmpl, err := template.ParseFS(blogTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "post.gohtml", post); err != nil {
		return err
	}

	return nil
}
