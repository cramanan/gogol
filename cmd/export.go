/*
Copyright © 2024
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use: "export",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("%s requires only one argument", cmd.Name())
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		err := RootDirectory.Read(args[0])
		if err != nil {
			return err
		}

		output, _ := cmd.Flags().GetString("output")
		expFile := NewFile(output + ".json")

		expFile.Content, err = json.Marshal(RootDirectory)

		if err != nil {
			return err
		}
		return expFile.Create(".")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringP("output", "o", "out", "select output file name")
}
