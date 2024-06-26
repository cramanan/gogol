/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

const GODEFAULT = `package main

import "fmt"

func main(){
	fmt.Println("Hello World")
}`

func GO(cmd *cobra.Command, args []string, root *Directory) (err error) {
	root.NewFile("main.go", []byte(GODEFAULT))
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Module name: ")
	mod, err := reader.ReadString('\n')
	root.NewFile("go.mod", []byte(
		fmt.Sprintf(
			"module %s\ngo %s",
			mod,
			strings.TrimPrefix(
				runtime.Version(),
				"go",
			),
		)))
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
