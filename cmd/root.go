/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var RootDirectory = NewDirectory("untitled")

const (
	GROUP_LANG = "LANG"

	FLAG_DOCKER   = "dockerfile"
	FLAG_ENV      = "env"
	FLAG_GITHUB   = "github"
	FLAG_LICENSE  = "license"
	FLAG_MAKEFILE = "makefile"
	FLAG_README   = "readme"
)

func init() {
	BoolP := rootCmd.PersistentFlags().BoolP
	// BoolP("all", "a", false, "all defaults files")
	BoolP(FLAG_DOCKER, "d", false, "add a Dockerfile")
	BoolP(FLAG_ENV, "e", false, "add a .env file")
	BoolP(FLAG_GITHUB, "g", false, "setup every .git files")
	BoolP(FLAG_LICENSE, "l", false, "add a LICENSE.md file")
	BoolP(FLAG_MAKEFILE, "m", false, "add a Makefile")
	BoolP(FLAG_README, "r", false, "add a README.md file")
	BoolP("tests", "t", false, "add a test directory")
}

var rootCmd = &cobra.Command{
	Use:   "gogol",
	Short: "Create projects faster than ever",
	Long: `
Generate simple projects directory structures
in the list of available languages.`,
	PersistentPreRunE:  PersistentPreRunE,
	PersistentPostRunE: PersistentPostRunE,
	CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
	SilenceUsage:       true,
}

func PersistentPreRunE(cmd *cobra.Command, args []string) error {
	if cmd.Name() == "help" {
		return nil
	}

	rootHasBoolFlag := cmd.Root().PersistentFlags().GetBool
	files := map[string]string{
		FLAG_DOCKER:   FLAG_DOCKER,
		FLAG_ENV:      ".env",
		FLAG_LICENSE:  "LICENSE.md",
		FLAG_MAKEFILE: FLAG_MAKEFILE,
		FLAG_README:   "README.md",
	}

	for flag := range files {
		value, _ := rootHasBoolFlag(flag)
		if value {
			RootDirectory.NewFile(files[flag])
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Project name: ")
	name, err := reader.ReadString('\n')
	if strings.Contains(name, "/") || strings.Contains(name, ".") {
		err = errors.New("project name cannot contain '/' or '.'")
	}
	if err != nil {
		return err
	}
	name = name[:len(name)-1]
	if name == "" {
		name = "untitled"
	}
	RootDirectory.Name = name

	github, _ := rootHasBoolFlag(FLAG_GITHUB)
	if github {
		RootDirectory.NewFile(".gitignore")
	}

	tests, _ := rootHasBoolFlag("tests")
	if tests {
		RootDirectory.NewDirectory("tests", &File{Name: ".gitkeep"})
		args := []*File{}
		if github {
			args = append(args, NewFile(".gitkeep"))
		}
		RootDirectory.NewDirectory("tests", args...)
	}

	return nil
}

func PersistentPostRunE(cmd *cobra.Command, args []string) (err error) {
	if cmd.Name() == "help" {
		return nil
	}

	if err = RootDirectory.Create("."); err != nil {
		return err
	}

	return
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
