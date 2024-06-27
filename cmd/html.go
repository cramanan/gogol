/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func HTML(cmd *cobra.Command, args []string, root *Directory) error {
	root.NewFile("index.html",
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
</html>`)
	root.NewFile("style.css", "*,\n*::before,\n*::after {\n\tmargin: 0;\n\tpadding: 0;\n\tbox-sizing: border-box;\n}\n\n")
	root.NewFile("script.js", "console.log('Hello World !')")
	return nil
}

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "",
	Long:  ``,
	Run:   GenerateFS(HTML),
}

func init() {
	rootCmd.AddCommand(htmlCmd)
}
