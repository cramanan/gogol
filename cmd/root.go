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

func init() {
	BoolP := rootCmd.PersistentFlags().BoolP
	BoolP("dockerfile", "d", false, "add a Dockerfile")
	BoolP("env", "e", false, "add a .env file")
	BoolP("github", "g", false, "setup every .git files")
	BoolP("license", "l", false, "add a LICENSE.md file")
	BoolP("makefile", "m", false, "add a Makefile")
	BoolP("readme", "r", false, "add a README.md file")
	BoolP("tests", "t", false, "add a test directory with a .gitkeep")
}

var rootCmd = &cobra.Command{
	Use:                "gogol",
	Short:              "Create a Go project",
	Long:               `Generate a simple Golang pro`,
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
		"dockerfile": "Dockerfile",
		"env":        ".env",
		"license":    "LICENSE.md",
		"makefile":   "Makefile",
		"readme":     "README.md",
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

	github, _ := rootHasBoolFlag("github")
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
