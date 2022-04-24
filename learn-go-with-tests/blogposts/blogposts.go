package blogposts

import (
	"io/fs"
)

type Post struct {
}

func FromFS(fsys fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(fsys, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
