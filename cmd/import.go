/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use: "import",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("%s requires only one argument", cmd.Name())
		}

		dt, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("error reading: %s", err.Error())
		}

		err = json.Unmarshal(dt, &RootDirectory)
		if err != nil {
			return fmt.Errorf("error unmarshaling : %s", err.Error())
		}

		return RootDirectory.Create(".")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
