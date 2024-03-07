/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"gogol/internal/tools"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

const GODEFAULT = `package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

`

func GetGolangVersion() (s string, err error) {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	s = string(out)
	return
}

func RunGo(cmd *cobra.Command, args []string) {
	fmt.Println(tools.OS())
	fmt.Print("Starting Golang Project...\n")
	_, versionErr := GetGolangVersion()
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
	}
	fmt.Print("Project name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		InternalError(err)
	}
	name = name[:len(name)-1]
	if name == "" {
		name = "Untitled"
	}
	fmt.Printf("Creating %s/ directory\n", name)
	err = Mkdir(name)
	if err != nil {
		InternalError(err)
	}

	fmt.Print("Package name: ")
	pkgname, err := reader.ReadString('\n')
	pkgname = pkgname[:len(pkgname)-1]
	if err != nil {
		InternalError(err)
	}
	if pkgname == "" {
		pkgname = "untitled"
	}

	if versionErr == nil {
		mod := exec.Command("go", "mod", "init", pkgname)
		if err := mod.Run(); err != nil {
			InternalError(err)
		}
	}

	main, err := os.Create("main.go")
	if err != nil {
		InternalError(err)
	}
	_, err = main.WriteString(GODEFAULT)
	if err != nil {
		InternalError(err)
	}
	fmt.Printf("All set and done !\nyou can now run:\n  cd %s\n  go run .\n", name)
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
