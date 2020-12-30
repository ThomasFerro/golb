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
	Title                string
	Locale               string
	PostPageTemplatePath string
	DistPath             string
}

type postData struct {
	Title   string
	Locale  string
	Content string
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

func generatePostsPages(metadata BlogMetadata, posts []posts.Post) (generatedPages, error) {
	splitTemplatePage := strings.Split(metadata.PostPageTemplatePath, "/")
	templateName := splitTemplatePage[len(splitTemplatePage)-1]
	postPageTemplate, err := template.New(templateName).ParseFiles(metadata.PostPageTemplatePath)
	if err != nil {
		return nil, fmt.Errorf("Cannot create the post page template: %w", err)
	}

	generatedPostsPages := generatedPages{}
	for _, post := range posts {
		data := postData{
			Title:   fmt.Sprintf("%v - %v", post.Name, metadata.Title),
			Locale:  metadata.Locale,
			Content: string(post.Content),
		}
		var bytesBuffer bytes.Buffer
		postPageWriter := bufio.NewWriter(&bytesBuffer)
		err := postPageTemplate.Execute(postPageWriter, data)
		if err != nil {
			return nil, err
		}
		generatedPostsPages = append(generatedPostsPages, generatedPage{
			content:  bytesBuffer.Bytes(),
			pagePath: fmt.Sprintf("posts/%v", formatPageName(post.Name)),
		})
	}
	return generatedPostsPages, nil
}

func createDirIfNeeded(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModePerm)
	}
}

func writeInDistFolder(distPath string, filesToWriteByType ...generatedPages) (GeneratedBlogPath, error) {
	createDirIfNeeded(distPath)
	for _, filesToWrite := range filesToWriteByType {
		for _, fileToWrite := range filesToWrite {
			filePath := fmt.Sprintf("%v/%v.html", distPath, fileToWrite.pagePath)

			dir := filepath.Dir(filePath)
			createDirIfNeeded(dir)

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
	// TODO: Generate the blog main page
	generatedPostsPages, err := generatePostsPages(metadata, posts)
	if err != nil {
		return "", fmt.Errorf("Cannot generate the post pages: %w", err)
	}

	return writeInDistFolder(metadata.DistPath, generatedPostsPages)
}
