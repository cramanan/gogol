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

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var RootDirectory = NewDirectory("untitled")

func init() {
	BoolP := RootCmd.PersistentFlags().BoolP
	BoolP("readme", "r", false, "add a README.md file")
	BoolP("license", "l", false, "add a LICENSE.md file")
	BoolP("dockerfile", "d", false, "add a Dockerfile")
	BoolP("makefile", "m", false, "add a Makefile")
	BoolP("github", "g", false, "initialize a github repo")
	BoolP("tests", "t", false, "add a test directory with a .gitkeep")
	BoolP("env", "e", false, "add a .env file")
}

var RootCmd = &cobra.Command{
	Use:                "gogol",
	PersistentPreRunE:  PersistentPreRunE,
	PersistentPostRunE: PersistentPostRunE,
	CompletionOptions:  cobra.CompletionOptions{DisableDefaultCmd: true},
	SilenceUsage:       true,
}

func PersistentPreRunE(cmd *cobra.Command, args []string) error {
	rootHasBoolFlag := cmd.Root().PersistentFlags().GetBool
	files := map[string]string{
		"readme":     "README.md",
		"license":    "LICENSE.md",
		"dockerfile": "Dockerfile",
		"makefile":   "Makefile",
		"env":        ".env",
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

	return nil
}

func PersistentPostRunE(cmd *cobra.Command, args []string) (err error) {
	err = RootDirectory.Create(".")
	if err != nil {
		return err
	}

	if github, _ := cmd.Root().PersistentFlags().GetBool("github"); github {
		_, err = git.PlainInit(RootDirectory.Name, false)
	}
	return
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
