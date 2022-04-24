package blogposts

import (
	"io/fs"
	"testing/fstest"
)

type Post struct {
}

func FromFS(fsys fstest.MapFS) []Post {
	dir, _ := fs.ReadDir(fsys, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}
