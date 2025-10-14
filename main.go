package main

import (
	"fmt"
	"log"

	"github.com/ThomasFerro/golb/config"

	"github.com/ThomasFerro/golb/blog"

	"github.com/ThomasFerro/golb/posts"
)

func main() {
	postsPath := config.GetConfiguration("POSTS_PATH")
	blogMetadata := blog.BlogMetadata{
		BlogTitle:            config.GetConfiguration("TITLE"),
		Locale:               config.GetConfiguration("LOCALE"),
		PostPageTemplatePath: config.GetConfiguration("POST_PAGE_TEMPLATE_PATH"),
		HomePageTemplatePath: config.GetConfiguration("HOME_PAGE_TEMPLATE_PATH"),
		DistPath:             config.GetConfiguration("DIST_PATH"),
		GlobalAssetsPath:     config.GetConfiguration("GLOBAL_ASSETS_PATH"),
	}
	blogPath, err := GenerateBlog(postsPath, blogMetadata)
	if err != nil {
		log.Fatalf("Unable to generate the blog: %v", err)
	}
	log.Printf("Blog successfully generated here: %v", blogPath)

}

func GenerateBlog(postsPath string, blogMetadata blog.BlogMetadata) (blog.GeneratedBlogPath, error) {
	postsExtractor := posts.NewFileSystemPostsExtractor(postsPath)
	extractedPosts, err := postsExtractor.ExtractPostsInformation()
	if err != nil {
		return "", fmt.Errorf("Unable to extract posts information: %w", err)
	}

	return blog.GenerateBlog(blogMetadata, extractedPosts)
}
