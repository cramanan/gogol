/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cramanan/gogol/internal/tools"

	"github.com/spf13/cobra"
)

func RunGo(cmd *cobra.Command, args []string) {
	fmt.Print("Starting Golang Project...\n")
	/*_, versionErr := GetGolangVersion()
	if versionErr != nil {
		fmt.Println("Golang is not installed or could not be found.\nTry running:\n  go version\nTo see if golang is installed")
		done := false
		for !done {
			choice, vErr := YesNo("Do you want to continue")
			if vErr != nil {
				fmt.Println(vErr)
				os.Exit(1)
			}
			switch strings.ToLower(choice) {
			case "y", "yes":
				done = true
			case "n", "no":
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please enter 'y' or 'n'")
			}
		}
	}*/
	fmt.Print("Project name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		InternalError(err)
	}
	name = name[:len(name)-1]
	if name == "" {
		name = "untitled"
	}

	fmt.Println("Fetching golang directory")
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
	pkgname, err := reader.ReadString('\n')
	if err != nil {
		InternalError(err)
	}
	pkgname = pkgname[:len(pkgname)-1]
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
	fmt.Printf("Creating %s/ directory\n", name)
	err = tools.CreateDirAndFiles(*dir)
	if err != nil {
		InternalError(err)
	}

	fmt.Println("All set and done !\nyou can now run:\n  go run .")
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
