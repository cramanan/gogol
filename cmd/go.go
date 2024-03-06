/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
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
	return string(out), nil
}

func YesNo(question string) (string, error) {
	fmt.Print(question + " [y/n]: ")
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return choice[:len(choice)-1], nil
}

func InternalError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func Run(cmd *cobra.Command, args []string) {
	fmt.Print("Starting Golang Project...\n")

	//TODO MUST CHECK IF GO IS INSTALLED
	_, err := GetGolangVersion()
	reader := bufio.NewReader(os.Stdin)
	if err != nil {
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
				fmt.Println("Invalid choice. Please enter 'y'/yes or 'n'")
			}
		}
	}
	fmt.Print("Project name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		InternalError(err)
	}
	name = name[:len(name)-1]
	if name == "" {
		name = "Untitled"
	}
	fmt.Printf("Creating %s/ directory\n", name)
	err = os.Mkdir(name, 0777)
	if err != nil {
		InternalError(err)
	}

	err = os.Chdir(name)
	if err != nil {
		InternalError(err)
	}

	fmt.Print("Package name: ")
	pkgname, err := reader.ReadString('\n')
	if err != nil {
		InternalError(err)
	}
	pkgname = pkgname[:len(pkgname)-1]
	mod := exec.Command("go", "mod", "init", pkgname)
	if err := mod.Run(); err != nil {
		InternalError(err)
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
	Long: `gogol go create a Golang project that includes the following files: 
	go.mod
	go.sum (optional)
	main.go
	main_test.go (optional)
	`,
	Run: Run,
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
