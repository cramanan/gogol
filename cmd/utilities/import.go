/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package utilities

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(importCmd)
}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use: "import",
	RunE: func(command *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("%s requires only one argument", command.Name())
		}

		dt, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("error reading: %s", err.Error())
		}

		err = json.Unmarshal(dt, &cmd.RootDirectory)
		if err != nil {
			return fmt.Errorf("error unmarshaling : %s", err.Error())
		}

		return cmd.RootDirectory.Create(".")
	},
}
