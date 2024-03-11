/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cramanan/gogol/internal/tools"

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
	fmt.Println("Starting Python project...")
	fmt.Print("Project name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		InternalError(err)
	}
	name = name[:len(name)-1]
	name = strings.ReplaceAll(name, "/", "")
	name = strings.ReplaceAll(name, ".", "")
	if name == "" {
		name = "untitled"
	}
	fmt.Printf("Creating %s/ directory\n", name)
	dir, err := tools.RetrieveYAMLdir("https://raw.githubusercontent.com/cramanan/gogol/cramanan/api/python.yaml")
	if err != nil {
		InternalError(err)
	}

	dir.Name = name
	if README {
		dir.AddFile(tools.File{Name: "README.md"})
	}

	if LICENSE {
		dir.AddFile(tools.File{Name: "LICENSE.md"})
	}

	f := dir.Search(fmt.Sprintf("%s/__main__.py", name))
	if f != nil {
		f.WriteString(tools.PYTHONDEFAULT)
	}

	err = dir.Create(".")
	if err != nil {
		InternalError(err)
	}

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
	│
	├── __main__.py
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
