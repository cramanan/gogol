/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

type Directory struct {
	Name        string
	repo        *git.Repository
	Directories map[string]*Directory
	Files       map[string]*File
}

func NewDirectory(name string) *Directory {
	return &Directory{
		name,
		nil,
		make(map[string]*Directory),
		make(map[string]*File),
	}
}

func NewFile(name string) *File {
	return &File{
		name,
		[]byte{},
	}
}

func (root *Directory) NewDirectory(name string, files ...*File) (d *Directory) {
	d = NewDirectory(name)
	for _, value := range files {
		d.NewFile(value.Name, value.Content)
	}
	root.Directories[name] = d
	return d
}

func (root *Directory) NewFile(name string, content ...[]byte) (f *File) {
	f = &File{
		name,
		[]byte{},
	}
	for _, b := range content {
		f.Content = append(f.Content, b...)
	}
	root.Files[name] = f
	return f
}

type File struct {
	Name    string
	Content []byte
}

var rootCmd = &cobra.Command{
	Use:   "gogol",
	Short: "",
	Long:  ``,
}

type CobraFunc func(cmd *cobra.Command, args []string, root *Directory) error

func GenerateFS(fn CobraFunc) func(cmd *cobra.Command, args []string) {
	rootHasBoolFlag := rootCmd.PersistentFlags().GetBool
	files := map[string]string{
		"readme":     "README.md",
		"license":    "LICENSE.md",
		"dockerfile": "Dockerfile",
		"makefile":   "Makefile",
	}

	return func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Project name: ")
		name, err := reader.ReadString('\n')
		if strings.Contains(name, "/") || strings.Contains(name, ".") {
			err = errors.New("project name cannot contain '/' or '.'")
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		name = name[:len(name)-1]
		if name == "" {
			name = "untitled"
		}

		root := NewDirectory(name)
		for flag := range files {
			value, _ := rootHasBoolFlag(flag)
			if value {
				root.NewFile(files[flag])
			}
		}

		tests, _ := rootHasBoolFlag("tests")
		github, _ := rootHasBoolFlag("github")

		var testsDir *Directory
		if tests {
			testsDir = root.NewDirectory("tests")
		}

		if github {
			root.NewFile(".gitignore")
			if testsDir != nil {
				testsDir.NewFile(".gitkeep")
			}
			defer func() {
				root.repo, err = git.PlainInit(root.Name, false)
				if err != nil {
					return
				}
			}()
		}

		fn(cmd, args, root)
		err = root.Create(".")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("All set and done ✓\nYou can now run:\n  cd %s\n", name)
	}
}

func (root Directory) Create(origin string) (err error) {
	var f func(string, Directory)
	f = func(path string, dir Directory) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return
		}
		for _, file := range dir.Files {
			ff, err := os.Create(filepath.Join(path, file.Name))
			if err != nil {
				return
			}
			defer ff.Close()
			_, err = ff.Write(file.Content)

			if err != nil {
				return
			}
		}
		for _, subdir := range dir.Directories {
			f(filepath.Join(path, subdir.Name), *subdir)
		}
	}
	f(filepath.Join(origin, root.Name), root)
	return err
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolP("readme", "r", false, "add a README.md file")
	rootCmd.PersistentFlags().BoolP("license", "l", false, "add a LICENSE.md file")
	rootCmd.PersistentFlags().BoolP("dockerfile", "d", false, "add a Dockerfile")
	rootCmd.PersistentFlags().BoolP("makefile", "m", false, "add a Makefile")
	rootCmd.PersistentFlags().BoolP("github", "g", false, "initialize a github repo")
	rootCmd.PersistentFlags().BoolP("tests", "t", false, "add a test directory with a .gitkeep")
}
