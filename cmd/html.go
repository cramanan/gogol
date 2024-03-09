/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cramanan/gogol/internal/tools"
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Create a simple HTML + CSS + JS Project",
	Long: `Create a directory containing the following file structure:
  directory/
    ├── index.html
    ├── script.js
    └── style.css
The index.html file will link the stylesheet & the script.
`,
	Run: RunHTML,
}

func RunHTML(cmd *cobra.Command, args []string) {
	fmt.Println("Starting an HTML/CSS/JS Project...")
	root, err := tools.RetrieveYAMLdir("https://raw.githubusercontent.com/cramanan/gogol/cramanan/api/html.yaml")
	if err != nil {
		InternalError(err)
	}
	fmt.Print("Project name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		InternalError(err)
	}
	name = name[:len(name)-1]
	if name == "" {
		name = "untitled"
	}
	fmt.Printf("Creating %s/ directory\n", name)
	root.Name = name
	f := root.Search(fmt.Sprintf("%s/index.html", root.Name))
	if f != nil {
		f.WriteString(tools.HTMLDEFAULT)
	}

	f = root.Search(fmt.Sprintf("%s/style.css", root.Name))
	if f != nil {
		f.WriteString(tools.CSSDEFAULT)
	}

	if err = tools.CreateDirAndFiles(*root); err != nil {
		InternalError(err)
	}
}

func init() {
	rootCmd.AddCommand(htmlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// htmlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// htmlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
