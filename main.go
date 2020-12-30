package main

import (
	"fmt"
	"log"

	"github.com/ThomasFerro/golb/posts"
)

func main() {
	// TODO: Config from env var
	postsExtractor := posts.NewFileSystemPostsExtractor("/Users/thomasferro/Documents/perso/git/readmes/posts")
	extractedPosts, err := postsExtractor.ExtractPostsInformation()
	if err != nil {
		log.Fatalf("Unable to extract posts information: %v", err)
	}
	fmt.Println(extractedPosts)
	// TODO: Generate the blog pages
}
