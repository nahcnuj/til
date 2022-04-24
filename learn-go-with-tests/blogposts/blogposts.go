package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title       string
	Description string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
)

func FromFS(fsys fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fsys, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPostFromFile(fsys, f)
		if err != nil {
			return nil, err //TODO: need clarification, totally fail if one file? or ignore?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPostFromFile(fsys fs.FS, f fs.DirEntry) (Post, error) {
	file, err := fsys.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	return newPost(file)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	title := readLine()[len(titlePrefix):]
	description := readLine()[len(descriptionPrefix):]

	post := Post{
		Title:       title,
		Description: description,
	}
	return post, nil
}
