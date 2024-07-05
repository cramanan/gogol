/*
Copyright © 2024
*/
package utilities

import (
	"encoding/json"
	"fmt"

	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use: "export",
	PreRunE: func(command *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("%s requires only one argument", command.Name())
		}
		return nil
	},
	RunE: func(command *cobra.Command, args []string) error {

		err := cmd.RootDirectory.Read(args[0])
		if err != nil {
			return err
		}

		output, _ := command.Flags().GetString("output")
		expFile := cmd.NewFile(fmt.Sprintf("%s.json", output))
		cmd.RootDirectory.Name = output

		b, err := json.MarshalIndent(cmd.RootDirectory, "", "\t")
		if err != nil {
			return err
		}

		expFile.Write(b)

		return expFile.Create(".")
	},
}

func init() {
	cmd.RootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringP("output", "o", "out", "select output file name")
}
