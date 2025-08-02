package cmd

import (
	"fmt"
	"gobook/pages"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var addr string
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve your book to a host",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		pages.Prebuild()

		for _, page := range pages.Site.Pages {
			func(p pages.Page) {
				http.HandleFunc("/"+p.URL, func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, pages.RenderPage(p))
				})

				http.HandleFunc("/"+p.URL+".md", func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, pages.RenderPage(p))
				})
			}(page)
		}

		http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/css")
			fmt.Fprintln(w, pages.Site.Stylesheet)
		})

		log.Println("Starting server on " + addr)
		http.ListenAndServe(addr, nil)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&addr, "addr", "a", "localhost:8080", "Address to use")
	rootCmd.AddCommand(serveCmd)
}
