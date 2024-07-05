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
		source.NewFile("main.rs", []byte("fn main() {\n\tprintln!(\"Hello, world!\");\n}\n"))
		cmd.RootDirectory.NewFile("Cargo.toml", []byte(`[package]
name = "untitled"
version = "0.1.0"
edition = "2021"

[dependencies]
`))
	},
	PostRunE: LanguagePostRunE,
}

func init() {
	cmd.RootCmd.AddCommand(rustCmd)
}
