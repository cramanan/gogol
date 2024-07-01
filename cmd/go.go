/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddGroup(&cobra.Group{
		ID:    GROUP_LANG,
		Title: "Languages",
	})
	goCmd.AddCommand(goWebCmd)
	rootCmd.AddCommand(goCmd)
}

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:       "go",
	GroupID:   GROUP_LANG,
	ValidArgs: []string{"web"},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Root().PersistentPreRunE(cmd, args)
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
			fmt.Sprintf("module %s\n\ngo %s\n",
				name,
				"1.19",
			),
		))
		RootDirectory.NewFile("main.go", []byte("package main"))

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		mainGo, ok := RootDirectory.Files["main.go"]
		if !ok {
			return errors.New("no main could be created")
		}
		mainGo.WriteString(`

import "fmt"
			
func main(){
	fmt.Println("Hello World")
}`)
		return nil
	},
}

var goWebCmd = &cobra.Command{
	Use: "web",
	RunE: func(cmd *cobra.Command, args []string) error {
		mainGo, ok := RootDirectory.Files["main.go"]
		if !ok {
			return errors.New("no main could be created")
		}
		mainGo.WriteString(`
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
}
`)

		api := RootDirectory.NewDirectory("api")
		api.NewDirectory("models")
		api.NewDirectory("controllers")

		static := RootDirectory.NewDirectory("static")
		static.NewDirectory("js")
		static.NewDirectory("css")
		return nil
	},
}
