/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
	"fmt"

	"github.com/cramanan/gogol/cmd"
	"github.com/cramanan/gogol/filesystem"
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	GroupID:           GROUP_LANG,
	Use:               "html",
	PersistentPreRunE: LanguagePreRunE,
	RunE: func(command *cobra.Command, args []string) error {
		index := filesystem.RootDirectory.NewFile("index.html")
		fmt.Fprintln(index, `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Document</title>
		<link rel="stylesheet" href="style.css" />
        <script src="script.js" defer ></script>
    </head>
    <body>
        <h1>Hello World</h1>
    </body>
</html>`)

		style := filesystem.RootDirectory.NewFile("style.css")
		fmt.Fprintln(style, "*,\n*::before,\n*::after {\n\tmargin: 0;\n\tpadding: 0;\n\tbox-sizing: border-box;\n}")

		script := filesystem.RootDirectory.NewFile("script.js")
		fmt.Fprintln(script, "console.log('Hello World !')")

		return nil
	},

	PersistentPostRunE: LanguagePostRunE,
}

func init() {
	cmd.RootCmd.AddCommand(htmlCmd)
}
