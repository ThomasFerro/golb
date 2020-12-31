package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

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

func generatePage(template *template.Template, pagePath string, data interface{}) (generatedPage, error) {
	bytesBuffer := new(bytes.Buffer)
	pageWriter := bufio.NewWriter(bytesBuffer)
	err := template.Execute(pageWriter, data)
	if err != nil {
		return generatedPage{}, err
	}
	pageWriter.Flush()
	return generatedPage{
		content:  bytesBuffer.Bytes(),
		pagePath: pagePath,
	}, nil
}

func getTemplate(templatePath string) (*template.Template, error) {
	splitTemplatePage := strings.Split(templatePath, "/")
	templateName := splitTemplatePage[len(splitTemplatePage)-1]
	return template.New(templateName).Funcs(template.FuncMap{
		"getPostPath": getPostPath,
	}).ParseFiles(templatePath)
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

func createPathToTheFileIfNeeded(filePath string) {
	dirPath := filepath.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
}

func writeInDistFolder(distPath string, filesToWriteByType ...generatedPages) (GeneratedBlogPath, error) {
	for _, filesToWrite := range filesToWriteByType {
		for _, fileToWrite := range filesToWrite {
			filePath := fmt.Sprintf("%v/%v.html", distPath, fileToWrite.pagePath)

			createPathToTheFileIfNeeded(filePath)

			file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				return "", err
			}
			defer file.Close()

			_, err = file.Write(fileToWrite.content)
			if err != nil {
				return "", err
			}
		}
	}
	return GeneratedBlogPath(distPath), nil
}

// GenerateBlog Generate a blog based on the posts and metadata
func GenerateBlog(metadata BlogMetadata, posts []posts.Post) (GeneratedBlogPath, error) {
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
