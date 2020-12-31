package main

import (
	"log"

	"github.com/ThomasFerro/golb/config"

	"github.com/ThomasFerro/golb/blog"

	"github.com/ThomasFerro/golb/posts"
)

func main() {
	postsExtractor := posts.NewFileSystemPostsExtractor(config.GetConfiguration("POSTS_PATH"))
	extractedPosts, err := postsExtractor.ExtractPostsInformation()
	if err != nil {
		log.Fatalf("Unable to extract posts information: %v", err)
	}

	blogPath, err := blog.GenerateBlog(blog.BlogMetadata{
		BlogTitle:            config.GetConfiguration("TITLE"),
		Locale:               config.GetConfiguration("LOCALE"),
		PostPageTemplatePath: config.GetConfiguration("POST_PAGE_TEMPLATE_PATH"),
		HomePageTemplatePath: config.GetConfiguration("HOME_PAGE_TEMPLATE_PATH"),
		DistPath:             config.GetConfiguration("DIST_PATH"),
		GlobalAssetsPath:     config.GetConfiguration("GLOBAL_ASSETS_PATH"),
	}, extractedPosts)
	if err != nil {
		log.Fatalf("Unable to generate the blog: %v", err)
	}
	log.Printf("Blog successfully generated here: %v", blogPath)
}
