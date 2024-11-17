/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cramanan/gogol/cmd"
	"github.com/cramanan/gogol/filesystem"
)

// pythonroot represents the python command
var pythonRoot = &cobra.Command{
	Use:     "python",
	GroupID: GROUP_LANG,
	PreRunE: LanguagePreRunE,
	Run: func(command *cobra.Command, args []string) {
		filesystem.RootDirectory.NewFile("__init__.py")
		main := filesystem.RootDirectory.NewFile("__main__.py")
		fmt.Fprintln(main, "print(\"Hello World\")")
		filesystem.RootDirectory.NewFile("setup.py")
		filesystem.RootDirectory.NewFile("requirements.txt")
	},
	PostRunE: LanguagePostRunE,
}

func init() {
	cmd.RootCmd.AddCommand(pythonRoot)
}
