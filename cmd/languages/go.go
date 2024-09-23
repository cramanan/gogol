/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:     "go",
	GroupID: GROUP_LANG,
	PersistentPreRunE: func(command *cobra.Command, args []string) error {
		err := LanguagePreRunE(command, args)
		if err != nil {
			return err
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Package name: ")
		name, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		if strings.Contains(name, "/") || strings.Contains(name, ".") {
			return errors.New("project name cannot contain '/' nor '.'")
		}

		name = name[:len(name)-1]
		if name == "" {
			name = "untitled"
		}

		mod := cmd.RootDirectory.NewFile("go.mod")
		fmt.Fprintf(mod, "module %s\n\ngo 1.19", name)

		return nil
	},

	RunE: func(command *cobra.Command, args []string) error {
		mainFile := cmd.RootDirectory.NewFile("main.go")
		_, err := fmt.Fprint(mainFile, "package main\n\nimport \"fmt\"\n\nfunc main(){\n\tfmt.Println(\"Hello World\")\n}")
		return err
	},

	PersistentPostRunE: LanguagePostRunE,
}
