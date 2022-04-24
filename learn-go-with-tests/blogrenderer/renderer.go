package blogrenderer

import (
	"blogposts"
	"html/template"
	"io"
)

const postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`

func Render(w io.Writer, post blogposts.Post) error {
	tmpl, err := template.New("post").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, post); err != nil {
		return err
	}

	return nil
}
