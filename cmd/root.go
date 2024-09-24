/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/cramanan/gogol/filesystem"
	"github.com/spf13/cobra"
)

var RootDirectory = filesystem.NewDirectory("untitled")

const (
	FLAG_DOCKER   = "dockerfile"
	FLAG_ENV      = "env"
	FLAG_GITHUB   = "github"
	FLAG_LICENSE  = "license"
	FLAG_MAKEFILE = "makefile"
	FLAG_README   = "readme"
)

var FILES_FLAGS = map[string]string{
	FLAG_DOCKER:   FLAG_DOCKER,
	FLAG_ENV:      ".env",
	FLAG_LICENSE:  "LICENSE.md",
	FLAG_MAKEFILE: FLAG_MAKEFILE,
	FLAG_README:   "README.md",
	// FLAG_GITHUB:   ".gitignore",
}

func init() {
	BoolP := RootCmd.PersistentFlags().BoolP

	BoolP(FLAG_DOCKER, "d", false, "add a Dockerfile")
	BoolP(FLAG_ENV, "e", false, "add a .env file")
	BoolP(FLAG_GITHUB, "g", false, "creates a .gitignore file")
	BoolP(FLAG_LICENSE, "l", false, "add a LICENSE.md file")
	BoolP(FLAG_MAKEFILE, "m", false, "add a Makefile")
	BoolP(FLAG_README, "r", false, "add a README.md file")
	BoolP("tests", "t", false, "add a test directory")
}

var RootCmd = &cobra.Command{
	Use:   "gogol",
	Short: "Create projects faster than ever",
	Long: `
Generate simple projects directory structures
in the list of available languages.`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	SilenceUsage:      true,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
