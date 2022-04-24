package blogrenderer

import (
	"blogposts"
	"bytes"
	"embed"
	"html/template"
	"io"

	"github.com/yuin/goldmark"
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
	vm, err := newPostVM(post)
	if err != nil {
		return err
	}
	return r.tmpl.ExecuteTemplate(w, "post.gohtml", vm)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	return r.tmpl.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	blogposts.Post
	HTMLBody template.HTML
}

func newPostVM(post blogposts.Post) (postViewModel, error) {
	buf := bytes.Buffer{}
	err := goldmark.Convert([]byte(post.Body), &buf)
	if err != nil {
		return postViewModel{}, err
	}

	vm := postViewModel{Post: post}
	vm.HTMLBody = template.HTML(buf.String())
	return vm, nil
}
