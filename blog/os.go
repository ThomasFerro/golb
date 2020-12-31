package blog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

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

func copyGlobalAssets(metadata BlogMetadata) error {
	if metadata.GlobalAssetsPath == "" {
		return nil
	}

	return filepath.Walk(metadata.GlobalAssetsPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		sourceFile, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		destinationFilePath := filepath.Join(metadata.DistPath, path)
		err = os.MkdirAll(filepath.Dir(destinationFilePath), os.ModePerm)
		if err != nil {
			return err
		}

		destinationFile, err := os.OpenFile(destinationFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return err
		}
		defer destinationFile.Close()

		_, err = io.Copy(destinationFile, sourceFile)
		if err != nil {
			return err
		}
		destinationFile.Sync()
		return nil
	})
}

func clearDist(destinationFolderPath string) error {
	return os.RemoveAll(destinationFolderPath)
}
