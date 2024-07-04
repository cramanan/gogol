/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:     "python",
	GroupID: GROUP_LANG,
	RunE: func(cmd *cobra.Command, args []string) error {
		RootDirectory.NewFile("__init__.py")
		RootDirectory.NewFile("__main__.py", []byte("print(\"Hello World\")"))
		RootDirectory.NewFile("setup.py")
		RootDirectory.NewFile("requirements.txt")
		RootDirectory.NewFile("__init__.py")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pythonCmd)
}
