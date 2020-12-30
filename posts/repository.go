package posts

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// TODO: Get assets (images)

// Repository A posts repository
type Repository interface {
	GetAllPosts() ([]Post, error)
}

type fileSystemPostsRepository struct {
	path string
}

func isTheFileOfAuthorizedExtension(authorizedExtensions []string, file os.FileInfo) bool {
	isFileOfAuthorizedExtension := false
	splitFileName := strings.Split(file.Name(), ".")
	fileExtension := splitFileName[len(splitFileName)-1]
	for _, extensionToCheck := range authorizedExtensions {
		if strings.EqualFold(fileExtension, extensionToCheck) {
			isFileOfAuthorizedExtension = true
		}
	}
	return isFileOfAuthorizedExtension
}

var markdownExtension = "md"

// TODO: Get the name and other metadata from the headers + fallback
func getPostName(fileContentAsMd string) string {
	r, _ := regexp.Compile(`(?m)^# (.*)$`)
	matchedString := r.FindStringSubmatch(fileContentAsMd)
	if len(matchedString) <= 1 {
		return ""
	}
	return matchedString[1]
}

func (repository fileSystemPostsRepository) GetAllPosts() ([]Post, error) {
	returnedPosts := []Post{}
	err := filepath.Walk(repository.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isTheFileOfAuthorizedExtension([]string{markdownExtension}, info) {
			fileContentAsMd, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			postName := getPostName(string(fileContentAsMd))
			returnedPosts = append(returnedPosts, Post{
				Name:      postName,
				Extension: markdownExtension,
				Content:   blackfriday.Run(fileContentAsMd),
			})
		}

		return nil
	})
	return returnedPosts, err
}

// NewFileSystemPostsRepository Creates a new posts repository based on the file system
func NewFileSystemPostsRepository(path string) Repository {
	return fileSystemPostsRepository{
		path,
	}
}
