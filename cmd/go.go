/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const GODEFAULT = `package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
`

func Run(cmd *cobra.Command, args []string) {
	fmt.Print("Starting Golang Project...\nProject name: ")

	//TODO MUST CHECK IF GO IS INSTALLED
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]
	if name == "" {
		name = "Untitled"
	}
	fmt.Printf("Creating %s/ directory\n", name)
	err := os.Mkdir(name, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	main, err := os.Create(name + "/main.go")
	_, err = main.WriteString(GODEFAULT)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Create a Golang project.",
	Long: `gogol go create a Golang project that includes the following files: 
	go.mod
	go.sum (optional)
	main.go
	main_test.go (optional)
	`,
	Run: Run,
}

func init() {
	rootCmd.AddCommand(goCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
