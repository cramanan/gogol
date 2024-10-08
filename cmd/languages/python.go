/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cramanan/gogol/cmd"
)

// pythonroot represents the python command
var pythonroot = &cobra.Command{
	Use:     "python",
	GroupID: GROUP_LANG,
	PreRunE: LanguagePreRunE,
	Run: func(command *cobra.Command, args []string) {
		cmd.RootDirectory.NewFile("__init__.py")
		main := cmd.RootDirectory.NewFile("__main__.py")
		fmt.Fprintln(main, "print(\"Hello World\")")
		cmd.RootDirectory.NewFile("setup.py")
		cmd.RootDirectory.NewFile("requirements.txt")
		cmd.RootDirectory.NewFile("__init__.py")
	},
	PostRunE: LanguagePostRunE,
}

func init() {
	cmd.RootCmd.AddCommand(pythonroot)
}
