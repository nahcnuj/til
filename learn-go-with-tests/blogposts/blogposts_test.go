package blogposts_test // appending "_test" to make tests closer to real usage

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	"blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("create posts from FS", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hello")},
			"hello-world2.md": {Data: []byte("コニチハ")},
		}

		posts, err := blogposts.FromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("return an error if ReadDir fails", func(t *testing.T) {
		_, err := blogposts.FromFS(StubFailingFS{})

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("something went wrong")
}
