package blogposts

import (
	"io/fs"
)

type Post struct {
	Title string
}

func FromFS(fsys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts, nil
}
