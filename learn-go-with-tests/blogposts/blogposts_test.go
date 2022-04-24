package blogposts_test // appending "_test" to make tests closer to real usage

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	t.Run("create posts from FS", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: TDD, go`
			secondBody = `Title: テスト投稿2
Description: ポストの概要
Tags: Perl, 正規表現`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.FromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		got := posts[0]
		want := blogposts.Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"TDD", "go"},
		}

		assertPost(t, got, want)
	})

	t.Run("return an error if ReadDir fails", func(t *testing.T) {
		_, err := blogposts.FromFS(StubFailingFS{})

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("something went wrong")
}
