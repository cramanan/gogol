/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	GroupID: GROUP_LANG,
	Use:     "html",
	Run: func(command *cobra.Command, args []string) {
		cmd.RootDirectory.NewFile("index.html").WriteString(`<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Document</title>
        <link rel="stylesheet" href="style.css" />
        <script src="script.js" defer />
    </head>
    <body>
        <h1>Hello World</h1>
    </body>
</html>`)
		cmd.RootDirectory.NewFile("style.css").WriteString("*,\n*::before,\n*::after {\n\tmargin: 0;\n\tpadding: 0;\n\tbox-sizing: border-box;\n}\n\n")
		cmd.RootDirectory.NewFile("script.js").WriteString("console.log('Hello World !')")
	},
}

func init() {
	cmd.RootCmd.AddCommand(htmlCmd)
}
