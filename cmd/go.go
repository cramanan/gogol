/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func GO(cmd *cobra.Command, args []string, root *Directory) (err error) {
	root.NewFile("main.go", []byte(
		`package main

import "fmt"

func main(){
	fmt.Println("Hello World")
}`))
	fmt.Print("Module name: ")
	reader := bufio.NewReader(os.Stdin)
	mod, err := reader.ReadString('\n')
	mod = mod[:len(mod)-1]
	if mod == "" {
		mod = root.Name
	}

	root.NewFile("go.mod",
		[]byte(fmt.Sprintf(
			"module %s\n\ngo %s\n",
			mod,
			runtime.Version()[2:], // Remove "go"
		)))
	return
}

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "go",
			Short: "Create a Golang project.",
			Long: `Generate a simple Golang project. 
The Go version specified in the generated go.mod file will match 
the version of the Go compiler used to build gogol.`,
			Run: GenerateFS(GO),
		},
	)
}
