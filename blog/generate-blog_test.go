package blog_test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/ThomasFerro/golb/blog"
	"github.com/ThomasFerro/golb/posts"
)

/*
TODO:
- Generate posts pages
- Copy global assets
*/

func TestGenerateTheHomePage(t *testing.T) {
	blogMetadata := blog.BlogMetadata{
		BlogTitle:            "My blog !",
		Locale:               "en",
		PostPageTemplatePath: "./postPageTemplate.go.html",
		HomePageTemplatePath: "./homePageTemplate.go.html",
		DistPath:             "../dist",
		GlobalAssetsPath:     "",
	}
	posts := []posts.Post{
		{
			Name:      "First post",
			Extension: ".md",
			Content:   []byte(""),
		},
		{
			Name:      "Second post",
			Extension: ".md",
			Content:   []byte(""),
		},
	}
	generatedBlogPath, err := blog.GenerateBlog(blogMetadata, posts)

	if err != nil {
		t.Fatalf("cannot generate the blog: %v", err)
	}

	expectedHomePage := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My blog !</title>
</head>
<body>
    <header><a href="/">My blog !</a></header>
    <ul class="posts">
        
            <li><a href="posts/second-post/">Second post</a></li>
        
            <li><a href="posts/first-post/">First post</a></li>
        
    </ul>
</body>
</html>`

	homePagePath := fmt.Sprintf("%v/index.html", generatedBlogPath)
	homePage, err := ioutil.ReadFile(homePagePath)
	if err != nil {
		t.Fatalf("cannot open the generated home page: %v", err)
	}

	if string(homePage) != expectedHomePage {
		t.Fatalf("the generated home page is not as expected, got: %v\nexpected: %v", string(homePage), expectedHomePage)
	}
}