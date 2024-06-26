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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gogol",
	Short: "",
	Long:  ``,
}

type CobraFunc func(cmd *cobra.Command, args []string, root *Directory) error

func GenerateFS(fn CobraFunc) func(cmd *cobra.Command, args []string) {
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
			name = "Untitled"
		}

		root := NewDirectory(name)
		files := map[string]string{
			"readme":     "README.md",
			"license":    "LICENSE.md",
			"dockerfile": "Dockerfile",
			"makefile":   "makefile",
		}

		for flag := range files {
			value, _ := rootCmd.PersistentFlags().GetBool(flag)
			if value {
				root.NewFile(files[flag])
			}
		}

		fn(cmd, args, root)
		err = root.Create(".")
		if err != nil {
			fmt.Println(err)
			return
		}

		github, _ := rootCmd.PersistentFlags().GetBool("github")
		if github {
			root.repo, err = git.PlainInit(root.Name, false)
			if err != nil {
				return
			}
		}
	}
}

func (root Directory) Create(origin string) (err error) {
	var f func(string, Directory)
	f = func(path string, dir Directory) {
		err = os.Mkdir(root.Name, os.ModePerm)
		if err != nil {
			return
		}
		for _, file := range dir.Files {
			ff, err := os.Create(filepath.Join(origin, path, file.Name))
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
	f(root.Name, root)
	return err
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().BoolP("readme", "r", false, "Help message for readme")
	rootCmd.PersistentFlags().BoolP("license", "l", false, "Help message for license")
	rootCmd.PersistentFlags().BoolP("dockerfile", "d", false, "Help message for license")
	rootCmd.PersistentFlags().BoolP("makefile", "m", false, "Help message for makefile")
	rootCmd.PersistentFlags().BoolP("github", "g", false, "Initialize a github repo")
}
