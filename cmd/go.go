/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cramanan/gogol/internal/tools"
	"github.com/go-git/go-git/v5"

	"github.com/spf13/cobra"
)

func RunGo(cmd *cobra.Command, args []string) {
	fmt.Print("Starting Golang Project...\n")
	fmt.Print("Project name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	name = name[:len(name)-1]
	if err != nil {
		InternalError(err)
	}
	name = strings.ReplaceAll(name, "/", "")
	name = strings.ReplaceAll(name, ".", "")

	fmt.Println("Fetching golang directory")
	dir, err := tools.RetrieveYAMLdir("https://raw.githubusercontent.com/cramanan/gogol/cramanan/api/golang.yaml")
	if err != nil {
		InternalError(err)
	}

	if name != "" {
		dir.Name = name
	}

	f := dir.Search(fmt.Sprintf("%s/main.go", dir.Name))
	if f != nil {
		f.WriteString(tools.GODEFAULT)
	}

	fmt.Print("Package name: ")
	pkgname, err := reader.ReadString('\n')
	pkgname = pkgname[:len(pkgname)-1]
	if err != nil {
		InternalError(err)
	}
	if pkgname == "" {
		pkgname = dir.Name
	}

	f = dir.Search(fmt.Sprintf("%s/go.mod", dir.Name))
	if f != nil {
		f.WriteString(fmt.Sprintf("module %s\n", pkgname))
	}

	if README {
		dir.AddFile(tools.File{Name: "README.md"})
	}

	if LICENSE {
		dir.AddFile(tools.File{Name: "LICENSE.md"})
	}

	if DOCKERFILE {
		dir.AddFile(tools.File{Name: "Dockerfile"})
	}

	if MAKEFILE {
		dir.AddFile(tools.File{Name: "Makefile"})
	}

	if GIT {
		dir.AddFile(tools.File{Name: ".gitignore"})
	}

	dir.PopFile(fmt.Sprintf("%s/go.sum", dir.Name))
	dir.PopFile(fmt.Sprintf("%s/main_test.go", dir.Name))
	fmt.Printf("Creating %s/ directory\n", dir.Name)
	err = dir.Create(".")
	if err != nil {
		InternalError(err)
	}

	if GIT {
		_, err = git.PlainInit(dir.Name, false)
		if err != nil {
			InternalError(err)
		}
	}

	fmt.Printf("All set and done !\nyou can now run:\n  cd %s\n  go run .\n", dir.Name)
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
