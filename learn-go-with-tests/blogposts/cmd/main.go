package main

import (
	"blogposts"
	"log"
	"os"
)

func main() {
	posts, err := blogposts.FromFS(os.DirFS("/tmp/posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
