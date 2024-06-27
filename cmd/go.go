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

const GODEFAULT = `package main

import "fmt"

func main(){
	fmt.Println("Hello World")
}`

func GO(cmd *cobra.Command, args []string, root *Directory) (err error) {
	root.NewFile("main.go", GODEFAULT)
	reader := bufio.NewReader(os.Stdin)
	mod, err := reader.ReadString('\n')
	mod = mod[:len(mod)-1]
	if mod == "" {
		mod = root.Name
	}

	root.NewFile("go.mod",
		fmt.Sprintf(
			"module %s\n\ngo %s\n",
			mod,
			runtime.Version()[2:], // Remove "go"
		))
	return
}

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "",
	Long:  ``,
	Run:   GenerateFS(GO),
}

func init() {
	rootCmd.AddCommand(goCmd)
}
