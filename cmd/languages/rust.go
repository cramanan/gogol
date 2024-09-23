/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
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
		source.NewFile("main.rs").WriteString("fn main() {\n\tprintln!(\"Hello, world!\");\n}\n")
		cmd.RootDirectory.NewFile("Cargo.toml").WriteString("[package]\nname = \"untitled\"\nversion = \"0.1.0\"\nedition = \"2021\"\n\n[dependencies]\n")
	},
	PostRunE: LanguagePostRunE,
}

func init() {
	cmd.RootCmd.AddCommand(rustCmd)
}
