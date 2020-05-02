package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const templateDir = "web/templates"
const outputDir = "dist"
const htmlBase = "web/base.html"

func main() {
	os.RemoveAll(outputDir)
	os.Mkdir(outputDir, 0755)

	htmlBaseTemp := template.Must(template.New("").ParseFiles(htmlBase))

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(templateDir, path)

		outPath := filepath.Join(outputDir, relPath)

		os.MkdirAll(filepath.Dir(outPath), 0755)

		file, _ := os.Open(path)
		defer file.Close()
		outFile, err := os.OpenFile(outPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		switch filepath.Ext(info.Name()) {
		case ".html":
			return renderHTML(htmlBaseTemp, file, outFile)
		default:
			return render(file, outFile)
		}
	})

	if err != nil {
		panic(err)
	}

	files, _ := ioutil.ReadDir(templateDir)
	for _, file := range files {
		if file.Name() == htmlBase {
			continue
		}
	}
}

func renderHTML(base *template.Template, page *os.File, w io.Writer) error {
	pageBytes, _ := ioutil.ReadAll(page)
	return base.ExecuteTemplate(w, "base.html", template.HTML(pageBytes))
}

func render(page *os.File, w io.Writer) error {
	_, err := io.Copy(w, page)
	return err
}
