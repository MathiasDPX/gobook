package pages

import (
	"embed"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

//go:embed template/*
var embeddedTemplates embed.FS

type Page struct {
	Title   string
	URL     string
	Content string
	Kwargs  map[string]string
}

var Site struct {
	Name       string `yaml:"name"`
	Template   string
	Sidebar    string
	Stylesheet string
	Pages      []Page
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

		if match == "pages\\INDEX.md" {
			page := Page{
				Title:   Site.Name,
				URL:     "",
				Content: filecontent,
			}

			pages = append(pages, page)
			continue
		}

		if match == "pages\\SUMMARY.md" {
			continue
		}

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
			Title:   result["title"] + " - " + Site.Name,
			URL:     strings.TrimSuffix(filepath.Base(match), ".md"),
			Content: content,
			Kwargs:  result,
		}

		pages = append(pages, page)
	}

	return pages
}

func RenderMarkdown(content string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(content))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	html := string(markdown.Render(doc, renderer))
	return html
}

func RenderPage(page Page) string {
	html := RenderMarkdown(page.Content)

	newTemplate := Site.Template
	newTemplate = strings.ReplaceAll(newTemplate, "{% title %}", page.Title)
	newTemplate = strings.ReplaceAll(newTemplate, "{% content %}", html)
	newTemplate = strings.ReplaceAll(newTemplate, "{% wiki_title %}", Site.Name)
	newTemplate = strings.ReplaceAll(newTemplate, "{% sidebar %}", Site.Sidebar)

	return newTemplate
}

func GetTemplateFS() (fs.FS, error) {
	path := filepath.Join(".", "template")

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return os.DirFS(path), nil
	}

	return fs.Sub(embeddedTemplates, "template")
}

func Prebuild() {
	// Load templates
	template_fs, err := GetTemplateFS()

	if err != nil {
		log.Fatal("Unable to get template filesystem")
	}

	raw_template, err := fs.ReadFile(template_fs, "index.html")

	if err != nil {
		log.Fatal("Unable to read template/index.html")
	} else {
		log.Println("Template loaded")
	}

	template := string(raw_template)

	// Load sidebar
	raw_sidebar, err := os.ReadFile("pages/SUMMARY.md")

	if err != nil {
		log.Println("Unable to load SUMMARY.md")
	}

	Site.Sidebar = RenderMarkdown(string(raw_sidebar))

	// Load site configuration
	siteData, err := os.ReadFile("_site.yml")
	if err != nil {
		log.Fatal("Unable to read _site.yml")
	}

	err = yaml.Unmarshal(siteData, &Site)
	if err != nil {
		log.Fatal("Unable to parse _site.yml")
	}

	Site.Template = template

	// Load pages
	pages := ProcessPages()
	Site.Pages = pages
	log.Printf("Loaded %d pages\n", len(pages))

	// Stylesheet
	content, err := fs.ReadFile(template_fs, "style.css")

	if err == nil {
		Site.Stylesheet = string(content)
	}
}
