package main_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	main "github.com/ThomasFerro/golb"
	"github.com/ThomasFerro/golb/blog"
)

func TestGenerateBlog(t *testing.T) {
	temporaryFolder, err := os.MkdirTemp("", "test_*")
	if err != nil {
		t.Fatalf("could not create temporary folder: %v", err)
	}
	defer func() {
		err := os.RemoveAll(temporaryFolder)
		if err != nil {
			log.Printf("could not remove temporary folder %v", err)
		}
	}()
	// TODO: Ajouter des posts, les templates et les assets
	postsPath := filepath.Join(temporaryFolder, "posts")
	err = os.Mkdir(postsPath, 0750)
	if err != nil {
		t.Fatalf("could not create posts folder: %v", err)
	}
	assetsPath := filepath.Join(temporaryFolder, "global-assets")
	err = os.Mkdir(assetsPath, 0750)
	if err != nil {
		t.Fatalf("could not create global-assets folder: %v", err)
	}
	distPath := filepath.Join(temporaryFolder, "dist")
	err = os.Mkdir(distPath, 0750)
	if err != nil {
		t.Fatalf("could not create dist folder: %v", err)
	}

	path, err := main.GenerateBlog(postsPath, blog.BlogMetadata{
		BlogTitle:            "Test blog",
		Locale:               "fr",
		PostPageTemplatePath: filepath.Join(temporaryFolder, "post-page-template.go.html"),
		HomePageTemplatePath: filepath.Join(temporaryFolder, "home-page-template.go.html"),
		DistPath:             distPath,
		GlobalAssetsPath:     assetsPath,
	})
	if err != nil {
		t.Fatalf("could not generate the blog: %v", err)
	}

	t.Errorf("not implemented: check generated path %v", path)
}
