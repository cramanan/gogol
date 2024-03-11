/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/cramanan/gogol/internal/tools"
	"github.com/go-git/go-git/v5"
	"golang.org/x/term"

	"github.com/spf13/cobra"
)

func RunGo(cmd *cobra.Command, args []string) {
	fmt.Print("Starting Golang Project...\r\n")
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	fmt.Print("Project name: ")
	t := term.NewTerminal(os.Stdin, "")
	name, err := t.ReadLine()
	if err != nil {
		InternalError(err)
	}
	name = strings.ReplaceAll(name, "/", "")
	name = strings.ReplaceAll(name, ".", "")
	if name == "" {
		name = "untitled"
	}

	fmt.Println("Fetching golang directory\r")
	dir, err := tools.RetrieveYAMLdir("https://raw.githubusercontent.com/cramanan/gogol/cramanan/api/golang.yaml")
	if err != nil {
		InternalError(err)
	}

	dir.Name = name
	f := dir.Search(fmt.Sprintf("%s/main.go", name))
	if f != nil {
		f.WriteString(tools.GODEFAULT)
	}

	fmt.Print("Package name: ")
	pkgname, err := t.ReadLine()
	if err != nil {

		InternalError(err)
	}
	if pkgname == "" {
		pkgname = "untitled"
	}

	f = dir.Search(fmt.Sprintf("%s/go.mod", name))
	if f != nil {
		f.Content = []byte(fmt.Sprintf("module %s\n", pkgname))
	}

	if README {
		dir.AddFile(tools.File{Name: "README.md"})
	}

	if LICENSE {
		dir.AddFile(tools.File{Name: "LICENSE.md"})
	}

	dir.PopFile(fmt.Sprintf("%s/go.sum", name))
	dir.PopFile(fmt.Sprintf("%s/main_test.go", name))
	fmt.Printf("Creating %s/ directory\r\n", name)
	err = dir.Create(".")
	if err != nil {
		InternalError(err)
	}

	if GIT {
		_, err := git.PlainInit(name, false)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("All set and done !\r\nyou can now run:\r\n  cd %s\r\n  go run .\r\n", dir.Name)
}

// goCmd represents the go command
var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Create a Golang project.",
	Long: `gogol go create a Golang project that includes the following file structure: 
  directory
    ├── go.mod
	├── go.sum (optional)
    ├── main_test.go (optional)
    └── main.go
	`,
	Run: RunGo,
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
