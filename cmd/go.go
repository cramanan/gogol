/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use: "go",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Root().PersistentPreRunE(cmd, args)
		RootDirectory.NewFile("main.go")
		if err != nil {
			return err
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Package name: ")
		name, err := reader.ReadString('\n')
		if strings.Contains(name, "/") || strings.Contains(name, ".") {
			err = errors.New("project name cannot contain '/' or '.'")
		}
		if err != nil {
			return err
		}
		name = name[:len(name)-1]
		if name == "" {
			name = "untitled"
		}
		RootDirectory.NewFile("go.mod", []byte(
			fmt.Sprintf("module %s\n\ngo %s",
				name,
				runtime.Version()[2:], // remove "go"
			),
		))
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
	},
}

var goWebCmd = &cobra.Command{
	Use: "web",
	RunE: func(cmd *cobra.Command, args []string) error {
		RootDirectory.NewDirectory("api")
		mainGo, ok := RootDirectory.Files["main.go"]
		if !ok {
			return errors.New("no main could be created")
		}
		mainGo.WriteString(
			`package main

import (
	"fmt"
	"net/http"
)
			
func main(){
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Printf("Server listening on port %s\n", server.Addr)
	server.ListenAndServe()
}`)
		return nil
	},
}

func init() {
	goCmd.AddCommand(goWebCmd)
	RootCmd.AddCommand(goCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
