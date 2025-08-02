package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initTmplCmd = &cobra.Command{
	Use:   "init-template [path]",
	Short: "Init a template for an existing ",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		fmt.Printf("Initializing a template in '%s'\n", path)

		os.MkdirAll(filepath.Join(path, "template"), os.ModePerm)
		err := ExtractFileFromEmbed("index.html", filepath.Join(path, "template", "index.html"))

		if err != nil {
			log.Println("Unable to copy file 'index.html")
			log.Fatal(err)
		}

		err = ExtractFileFromEmbed("style.css", filepath.Join(path, "template", "style.css"))

		if err != nil {
			log.Println("Unable to copy file 'style.css")
			log.Fatal(err)
		}

		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(initTmplCmd)
}
