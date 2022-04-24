package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
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

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titlePrefix)
	description := readMetaLine(descriptionPrefix)
	tagsLine := readMetaLine(tagsPrefix)
	tags := strings.Split(tagsLine, ", ")

	body := readBody(scanner)

	post := Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // discard a "---" line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func (p Post) SanitizedTitleForUrl() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
