/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func HTML(cmd *cobra.Command, args []string, root *Directory) error {
	root.NewFile("index.html", []byte(
		`<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Document</title>
        <link rel="stylesheet" href="style.css" />
        <script src="script.js" defer></script>
    </head>
    <body>
        <h1>Hello World</h1>
    </body>
</html>`))
	root.NewFile("style.css", []byte("*,\n*::before,\n*::after {\n\tmargin: 0;\n\tpadding: 0;\n\tbox-sizing: border-box;\n}\n\n"))
	root.NewFile("script.js", []byte("console.log('Hello World !')"))
	return nil
}

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "html",
			Short: "Create an HTML project.",
			Long: `Generate an HTML, CSS, and JavaScript project.

The HTML header will reference the associated CSS and JS files.`,
			Run: GenerateFS(HTML),
		})
}
