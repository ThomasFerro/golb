package posts

import (
	"fmt"
)

// Extractor Service used to extract posts information
type Extractor interface {
	ExtractPostsInformation() ([]Post, error)
}

func processBlogPost(post Post) (*Post, error) {
	// TODO: Process md files => get header (shouldBePublished) then convert to HTML (?)
	return &post, nil
}

func processBlogPosts(posts []Post) ([]Post, error) {
	processedBlogPosts := []Post{}

	for _, nextPostToProcess := range posts {
		processedBlogPost, err := processBlogPost(nextPostToProcess)
		if err != nil {
			return nil, err
		}
		if processedBlogPost != nil {
			processedBlogPosts = append(processedBlogPosts, *processedBlogPost)
		}
	}

	return processedBlogPosts, nil
}

type fileSystemPostsExtractor struct {
	repository Repository
}

// ExtractPostsInformation Extract the posts information for the blog
func (postsExtractor fileSystemPostsExtractor) ExtractPostsInformation() ([]Post, error) {
	blogPosts, err := postsExtractor.repository.GetAllPosts()
	if err != nil {
		return nil, fmt.Errorf("Unable to get all posts from the repository: %w", err)
	}
	return processBlogPosts(blogPosts)
}

// NewFileSystemPostsExtractor Creates a new FileSystemPostsExtractor
func NewFileSystemPostsExtractor(path string) Extractor {
	return fileSystemPostsExtractor{
		repository: NewFileSystemPostsRepository(path),
	}
}
