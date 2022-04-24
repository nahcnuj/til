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

type PostRenderer struct {
	tmpl *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	tmpl, err := template.ParseFS(blogTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{tmpl}, nil
}

func (r *PostRenderer) Render(w io.Writer, post blogposts.Post) error {
	if err := r.tmpl.ExecuteTemplate(w, "post.gohtml", post); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitizedTitleForUrl}}">{{.Title}}</a></li>{{end}}</ol>`

	tmpl, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, posts); err != nil {
		return err
	}
	return nil
}
