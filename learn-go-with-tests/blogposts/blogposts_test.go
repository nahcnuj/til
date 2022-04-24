package blogposts_test // appending "_test" to make tests closer to real usage

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hello")},
		"hello-world2.md": {Data: []byte("コニチハ")},
	}

	posts := blogposts.FromFS(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}
