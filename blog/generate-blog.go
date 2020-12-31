package blog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ThomasFerro/golb/posts"
)

// GeneratedBlogPath The path to the generated blog
type GeneratedBlogPath string

type BlogMetadata struct {
	BlogTitle            string
	Locale               string
	PostPageTemplatePath string
	HomePageTemplatePath string
	DistPath             string
	GlobalAssetsPath     string
}

type homeData struct {
	BlogTitle string
	Locale    string
	Posts     []posts.Post
}

type postData struct {
	BlogTitle string
	PageTitle string
	Locale    string
	Content   string
}

type generatedPage struct {
	pagePath string
	content  []byte
}
type generatedPages []generatedPage

func formatPageName(pageName string) string {
	replaceSpacesWithDashes := regexp.MustCompile(` `)
	processedPageName := replaceSpacesWithDashes.ReplaceAllString(pageName, "-")
	regex := regexp.MustCompile("[^a-zA-Z0-9--]+")
	processedPageName = regex.ReplaceAllString(processedPageName, "")
	return strings.ToLower(processedPageName)
}

func getPostPath(post posts.Post) string {
	return fmt.Sprintf("posts/%v/", formatPageName(post.Name))
}

func generatePostsPages(metadata BlogMetadata, posts []posts.Post) (generatedPages, error) {
	postPageTemplate, err := getTemplate(metadata.PostPageTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot create the post page template: %w", err)
	}

	generatedPostsPages := generatedPages{}
	for _, post := range posts {
		data := postData{
			BlogTitle: metadata.BlogTitle,
			PageTitle: fmt.Sprintf("%v - %v", post.Name, metadata.BlogTitle),
			Locale:    metadata.Locale,
			Content:   string(post.Content),
		}
		postPagePath := fmt.Sprintf("%vindex", getPostPath(post))
		generatedPostPage, err := generatePage(postPageTemplate, postPagePath, data)
		if err != nil {
			return nil, err
		}
		generatedPostsPages = append(generatedPostsPages, generatedPostPage)
	}
	return generatedPostsPages, nil
}

func reversePosts(postsToReverse []posts.Post) []posts.Post {
	reversedPosts := []posts.Post{}
	for postIndex := len(postsToReverse) - 1; postIndex >= 0; postIndex-- {
		reversedPosts = append(reversedPosts, postsToReverse[postIndex])
	}
	return reversedPosts
}

func generateHomePage(metadata BlogMetadata, posts []posts.Post) (generatedPages, error) {
	homePageTemplate, err := getTemplate(metadata.HomePageTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot create the home page template: %w", err)
	}
	data := homeData{
		BlogTitle: metadata.BlogTitle,
		Locale:    metadata.Locale,
		Posts:     reversePosts(posts),
	}
	generatedHomePage, err := generatePage(homePageTemplate, "index", data)
	if err != nil {
		return nil, err
	}

	return generatedPages{
		generatedHomePage,
	}, nil
}

// GenerateBlog Generate a blog based on the posts and metadata
func GenerateBlog(metadata BlogMetadata, posts []posts.Post) (GeneratedBlogPath, error) {
	err := clearDist(metadata.DistPath)
	if err != nil {
		return "", fmt.Errorf("Cannot clear destination folder: %w", err)
	}

	err = copyGlobalAssets(metadata)
	if err != nil {
		return "", fmt.Errorf("Cannot copy global assets: %w", err)
	}

	generatedHomePage, err := generateHomePage(metadata, posts)
	if err != nil {
		return "", fmt.Errorf("Cannot generate the homepage: %w", err)
	}
	generatedPostsPages, err := generatePostsPages(metadata, posts)
	if err != nil {
		return "", fmt.Errorf("Cannot generate the post pages: %w", err)
	}

	return writeInDistFolder(metadata.DistPath, generatedHomePage, generatedPostsPages)
}
