/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
	"fmt"

	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

// rustCmd represents the rust command
var rustCmd = &cobra.Command{
	Use:     "rust",
	GroupID: GROUP_LANG,
	PreRunE: LanguagePreRunE,
	Run: func(command *cobra.Command, args []string) {
		source := cmd.RootDirectory.NewDirectory("src")
		main := source.NewFile("main.rs")
		fmt.Fprintln(main, "fn main() {\n\tprintln!(\"Hello, world!\");\n}")
		lockFile := cmd.RootDirectory.NewFile("Cargo.toml")
		fmt.Fprintln(lockFile, "[package]\nname = \"untitled\"\nversion = \"0.1.0\"\nedition = \"2021\"\n\n[dependencies]")
	},
	PostRunE: LanguagePostRunE,
}

func init() {
	cmd.RootCmd.AddCommand(rustCmd)
}
