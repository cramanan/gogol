/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package languages

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cramanan/gogol/cmd"
	"github.com/spf13/cobra"
)

const GROUP_LANG = "LANG"

func init() {
	cmd.RootCmd.AddGroup(&cobra.Group{
		ID:    GROUP_LANG,
		Title: "Languages",
	})
	goCmd.AddCommand(goWebCmd)
	cmd.RootCmd.AddCommand(goCmd)
}

func LanguagePreRunE(command *cobra.Command, args []string) error {
	rootHasBoolFlag := command.PersistentFlags().GetBool
	files := map[string]string{
		cmd.FLAG_DOCKER:   cmd.FLAG_DOCKER,
		cmd.FLAG_ENV:      ".env",
		cmd.FLAG_LICENSE:  "LICENSE.md",
		cmd.FLAG_MAKEFILE: cmd.FLAG_MAKEFILE,
		cmd.FLAG_README:   "README.md",
	}

	for flag := range files {
		value, _ := rootHasBoolFlag(flag)
		if value {
			cmd.RootDirectory.NewFile(files[flag])
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Project name: ")
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
	cmd.RootDirectory.Name = name

	github, _ := rootHasBoolFlag(cmd.FLAG_GITHUB)
	if github {
		cmd.RootDirectory.NewFile(".gitignore")
	}

	tests, _ := rootHasBoolFlag("tests")
	if tests {
		cmd.RootDirectory.NewDirectory("tests")
		args := []*cmd.File{}
		if github {
			args = append(args, cmd.NewFile(".gitkeep"))
		}
		cmd.RootDirectory.NewDirectory("tests", args...)
	}

	return nil
}

func LanguagePostRunE(*cobra.Command, []string) (err error) {
	return cmd.RootDirectory.Create(".")
}

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:       "go",
	GroupID:   GROUP_LANG,
	ValidArgs: []string{"web"},
	PersistentPreRunE: func(command *cobra.Command, args []string) error {
		err := LanguagePreRunE(command, args)
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
		cmd.RootDirectory.NewFile("go.mod",
			fmt.Sprintf("module %s\n\ngo %s\n",
				name,
				"1.19",
			))
		cmd.RootDirectory.NewFile("main.go", ("package main"))

		return nil
	},

	RunE: func(command *cobra.Command, args []string) error {
		mainGo, ok := cmd.RootDirectory.Files["main.go"]
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
	PersistentPostRunE: LanguagePostRunE,
}

var goWebCmd = &cobra.Command{
	Use: "web",
	Run: func(command *cobra.Command, args []string) {
		mainGo, ok := cmd.RootDirectory.Files["main.go"]
		if !ok {
			mainGo = cmd.RootDirectory.NewFile("main.go")
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

		api := cmd.RootDirectory.NewDirectory("api")
		api.NewFile("api.go", ("package api"))
		for _, name := range []string{"models", "controllers"} {
			directory := api.NewDirectory(name)
			directory.NewFile(fmt.Sprintf("%s.go", name), (fmt.Sprintf("package %s", name)))
		}

		static := cmd.RootDirectory.NewDirectory("static")
		for _, name := range []string{"css", "js"} {
			static.NewDirectory(name)
		}
	},
}
