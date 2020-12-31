package blog

import (
	"bufio"
	"bytes"
	"strings"
	"text/template"
)

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
