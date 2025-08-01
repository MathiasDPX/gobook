package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Page struct {
	title   string
	url     string
	content string
	kwargs  map[string]string
}

var Site struct {
	Name     string `yaml:"name"`
	Template string
}

func ProcessPages() []Page {
	matches, err := filepath.Glob("pages/*md")

	if err != nil {
		log.Fatal(err)
	}

	var pages = []Page{}

	for _, match := range matches {
		rawcontent, err := os.ReadFile(match)
		if err != nil {
			continue
		}
		filecontent := string(rawcontent)

		parts := strings.SplitN(filecontent, "---", 3)
		headers := parts[1]
		headers = strings.TrimSpace(headers)

		content := parts[2]
		content = strings.TrimSpace(content)

		lines := strings.Split(headers, "\n")
		result := make(map[string]string)

		for _, line := range lines {
			part := strings.SplitN(line, ":", 2)
			if len(part) == 2 {
				result[part[0]] = strings.TrimSpace(part[1])
			}
		}

		page := Page{
			title:   result["title"],
			url:     strings.TrimSuffix(filepath.Base(match), ".md"),
			content: content,
			kwargs:  result,
		}

		pages = append(pages, page)
	}

	return pages
}

func RenderPage(page Page) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(page.content))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	html := string(markdown.Render(doc, renderer))

	newTemplate := Site.Template
	newTemplate = strings.Replace(newTemplate, "{% title %}", page.title, -1)
	newTemplate = strings.Replace(newTemplate, "{% content %}", html, -1)
	newTemplate = strings.Replace(newTemplate, "{% sidebar %}", "", -1)
	newTemplate = strings.Replace(newTemplate, "{% wiki_title %}", Site.Name, -1)

	return newTemplate
}

func main() {
	// Load templates
	raw_template, err := os.ReadFile("template/index.html")

	if err != nil {
		log.Fatal("Unable to read template/index.html")
	} else {
		log.Println("Template loaded")
	}

	template := string(raw_template)

	// Load pages
	pages := ProcessPages()
	log.Printf("Loaded %d pages\n", len(pages))
	for _, page := range pages {
		func(p Page) {
			http.HandleFunc("/"+p.url, func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, RenderPage(p))
			})
		}(page)
	}

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		content, _ := os.ReadFile("template/style.css")

		w.Header().Set("Content-Type", "text/css")
		fmt.Fprintln(w, string(content))
	})

	// Load site configuration
	siteData, err := os.ReadFile("_site.yml")
	if err != nil {
		log.Fatal("Unable to read _site.yml:", err)
	}

	err = yaml.Unmarshal(siteData, &Site)
	if err != nil {
		log.Fatal("Unable to parse _site.yml:", err)
	}

	Site.Template = template

	log.Println("Starting server on port 8080")
	http.ListenAndServe("localhost:8080", nil)
}
