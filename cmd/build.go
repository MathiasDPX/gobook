package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/MathiasDPX/gobook/pages"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build your book to static html/css",
	Run: func(cmd *cobra.Command, args []string) {
		pages.Prebuild()
		os.MkdirAll("_book", os.ModePerm)

		for _, page := range pages.Site.Pages {
			WriteFile(filepath.Join("_book", pages.GetHTMLFileName(page)), pages.RenderPage(page))
			log.Printf("Built '%s' page\n", page.URL)
		}

		WriteFile(filepath.Join("_book", "style.css"), pages.Site.Stylesheet)
		log.Printf("Built 'style.css'")

		log.Println("Site built in _book")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
