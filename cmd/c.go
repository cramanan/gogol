/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RunC(cmd *cobra.Command, args []string) {
	fmt.Println("c called")
}

// cCmd represents the c command
var cCmd = &cobra.Command{
	Use:   "c",
	Short: "Create a C project.",
	Long: `gogol c create a simple c project with the following file structure:
  directory/
    ├── makefile (optional)
    └── src
        ├── header.h (optional)
        └── main.c
`,
	Run: RunC,
}

func init() {
	rootCmd.AddCommand(cCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
