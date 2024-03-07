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

	"github.com/spf13/cobra"
)

const CDEFAULT = `#include <stdio.h>

int main(int argc, char *argv[]){
	printf("Hello World\n");
    return 0;
}

`

func GetGCCVersion() (s string, err error) {
	cmd := exec.Command("gcc", "--version")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	s = string(out)
	return
}

func RunC(cmd *cobra.Command, args []string) {
	fmt.Println("Startin C project...")
	_, err := GetGCCVersion()
	if err != nil {
		fmt.Println("GCC is not installed or could not be found.\nTry running:\n  gcc --version\nTo see if gcc is installed")
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
	fmt.Print("Creating source folder...")
	err = Mkdir("src")
	if err != nil {
		InternalError(err)
	}
	fmt.Println("done")
	main, err := os.Create("main.c")
	if err != nil {
		InternalError(err)
	}
	_, err = main.WriteString(CDEFAULT)
	if err != nil {
		InternalError(err)
	}
	fmt.Printf("All set and done !\nyou can now run:\n  cd %s\n  gcc -o <output_name> src/main.c\n ./<output_name>\n", name)
}

// cCmd represents the c command
var cCmd = &cobra.Command{
	Use:   "c",
	Short: "Create a C project.",
	Long: `gogol c create a simple c project with the following file structure:
  directory/
    ├── makefile (optional)
    └── src
        ├── header.h (optional)
        └── main.c
`,
	Run: RunC,
}

func init() {
	rootCmd.AddCommand(cCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
