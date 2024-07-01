/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// htmlCmd represents the html command
var htmlCmd = &cobra.Command{
	GroupID: GROUP_LANG,
	Use:     "html",
	Run: func(cmd *cobra.Command, args []string) {
		RootDirectory.NewFile("index.html", []byte(
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
		RootDirectory.NewFile("style.css", []byte("*,\n*::before,\n*::after {\n\tmargin: 0;\n\tpadding: 0;\n\tbox-sizing: border-box;\n}\n\n"))
		RootDirectory.NewFile("script.js", []byte("console.log('Hello World !')"))
	},
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
