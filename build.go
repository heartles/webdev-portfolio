package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const templateDir = "web/templates"
const outputDir = "dist"
const htmlBase = "web/base.html"

// template rules:
// 1. Parse HTML files and inject them into the base template UNLESS
// 2. If the HTML file ends in .raw.html, then copy it without
//    any templating
// 3. Copy all other files without any templating (will change later if needed)

func main() {
	os.RemoveAll(outputDir)
	os.Mkdir(outputDir, 0755)

	htmlBaseTemp := template.Must(template.New("").ParseFiles(htmlBase))

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() {
			return nil
		}

		relPath, _ := filepath.Rel(templateDir, path)

		outPath := strings.Replace(filepath.Join(outputDir, relPath), ".raw.", ".", 1)

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
			if !strings.HasSuffix(info.Name(), ".raw.html") {
				return renderHTML(htmlBaseTemp, outFile, file)
			}

			fallthrough
		default:
			_, err := io.Copy(outFile, file)
			return err
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

func renderHTML(base *template.Template, out io.Writer, page io.Reader) error {
	pageBytes, _ := ioutil.ReadAll(page)
	return base.ExecuteTemplate(out, "base.html", template.HTML(pageBytes))
}
