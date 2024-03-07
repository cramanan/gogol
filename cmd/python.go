/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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

func GetPythonVersion() (s string, err error) {
	cmd := exec.Command("python3", "--version")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	s = string(out)
	return
}

func RunPython(cmd *cobra.Command, args []string) {
	fmt.Println("Startin Python project...")
	_, err := GetPythonVersion()
	if err != nil {
		fmt.Println("Python is not installed or could not be found.\nTry running:\n  python3 --version\nTo see if python is installed")
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
	dir, err := tools.RetrieveYAMLdir("../api/python.yaml")
	if err != nil {
		InternalError(err)
	}
	fmt.Println(dir)
	fmt.Printf("All set and done !\nyou can now run:\n  cd %s\n  python3 .\n", name)
}

// pythonCmd represents the python command
var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Create a python project.",
	Long: `gogol python create a python project with the following file structure:
  directory/
	├── main/
	│   ├── __init__.py
	│   └── helpers.py
	│
	├── tests/
	│	├── main_tests.py
	│	└── helpers_tests.py
	├── requirements.txt
	└── setup.py
    `,
	Run: RunPython,
}

func init() {
	rootCmd.AddCommand(pythonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pythonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pythonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
