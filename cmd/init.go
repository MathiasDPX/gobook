package cmd

import (
	"bufio"
	"fmt"
	"gobook/pages"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func WriteFile(path string, content string) bool {
	data := []byte(content)
	err := os.WriteFile(path, data, 0644)

	return err != nil
}

func CopyFile(source string, destination string) error {
	filesys, err := pages.GetTemplateFS()

	if err != nil {
		return err
	}

	data, err := fs.ReadFile(filesys, source)

	if err != nil {
		return err
	}

	err = os.WriteFile(destination, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Init a GoBook book",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		fmt.Printf("Initializing a book in '%s'\n", path)

		fmt.Print("Book name: ")
		name, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		name = strings.TrimSpace(name)

		os.MkdirAll(filepath.Join(path, "pages"), os.ModePerm)
		WriteFile(filepath.Join(path, "_site.yml"), "name: "+name)
		WriteFile(filepath.Join(path, "pages", "INDEX.md"), "# "+name+"\n\nWelcome to your brand new GoBook book")

		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
